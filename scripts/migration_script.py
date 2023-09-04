import os
import glob
import logging, sys

'''
For terraform-provider-opal migration to v2.0.0, we need to wrap certain fields in the resource block with a new block.

This script will take a terraform file and wrap the fields in the resource block with a new block called "request_configuration".

Currently, we only support parsing files in which opal_resource and opal_group blocks do NOT contain:
- multiline comments
- multiline strings

While we handle this case in the script, it is recommended that the reviewer_stage blocks are properly formatted, 
i.e. the opening brackets are on the same line as the block name and the closing brackets are on their own line.

e.g.
resource "opal_resource" "my_resource" {
  reviewer_stage {
    require_manager_approval = true 
  }

  reviewer_stage {
    reviewer {
      id = opal_owner.security.id
    }

    reviewer {
      id = opal_owner.data.id
    }
  }
}

and not:
resource "opal_resource" "complex_reviewer_config" {
  // ...

  // Here first the manager has to approve. Once the manager has approved, both the security owner and the data owner need to approve
  // NOTE: The ordering determines the ordering of the stages
  reviewer_stage { require_manager_approval = true }

  reviewer_stage {
    reviewer { id = opal_owner.security.id }

    reviewer {
      id = opal_owner.data.id
    }
  }
}
'''

logging.basicConfig(stream=sys.stderr, level=logging.INFO)

RESOURCE_TYPES = ["opal_resource", "opal_group"]
FIELDS_TO_WRAP = ["auto_approval", "require_mfa_to_request", "require_support_ticket", "max_duration", "recommended_duration", "request_template_id", "is_requestable"]
WRAPPED_BLOCK_NAME = "request_configuration"
NESTED_BLOCK_NAMES = ["reviewer_stage"]
TAB = " " * 2
COMMENT = "//"

def get_indentation(line):
    return len(line) - len(line.lstrip())

def is_resource_line(line):
    return any(line.strip().startswith(f'resource "{res_type}"') for res_type in RESOURCE_TYPES)

def is_field_to_wrap(line):
    return any(line.strip().startswith(f"{field} =") for field in FIELDS_TO_WRAP)

def is_start_of_nested_block(line):
    return any(line.strip().startswith(f"{block_name} {{") for block_name in NESTED_BLOCK_NAMES)

def count_bracket_changes(line):
    bracket_count = 0
    is_escaped = False
    is_in_string = False
    is_in_comment = False
    for char in line:
        if is_escaped:
            is_escaped = False
            continue
        if char == "\\":
            is_escaped = True
            continue
        if char == "\"":
            is_in_string = not is_in_string
            continue
        if char == COMMENT:
            is_in_comment = True
            continue
        if is_in_string:
            continue
        if is_in_comment:
            if char == "\n":
                is_in_comment = False
            continue
        if char == "{":
            logging.debug(f"Found opening bracket: {line}")
            bracket_count += 1
        elif char == "}":
            logging.debug(f"Found closing bracket: {line}")
            bracket_count -= 1
    return bracket_count
            

def extract_nested_block(lines, indentation_level):
    nested_block = []
    bracket_count = 1
    indent = " " * indentation_level

    while bracket_count > 0:
        line = next(lines)
        stripped_line = line.strip()

        if stripped_line.startswith(COMMENT):
            nested_block.append(line)
            continue
        
        if count_bracket_changes(stripped_line) != 0:
            bracket_count += count_bracket_changes(stripped_line)

        nested_block.append(indent + line)
    return nested_block

def modify_terraform_with_nested_blocks(s):
    lines = iter(s.split("\n"))  # Convert lines to iterator
    in_opal_resource = False
    new_lines = []
    wrapped_block_data = []
    nested_blocks = []
    indentation_level = 0
    bracket_count = 0  # To keep track of nested blocks

    for line in lines:
        stripped_line = line.strip()

        if is_resource_line(stripped_line):
            logging.info(f"Found resource line: {stripped_line}")
            in_opal_resource = True
            bracket_count = 1  # Opening bracket of the resource block
            indentation_level = get_indentation(line) + 2  # 2 spaces for indentation within the resource block
            new_lines.append(line)
            continue

        if in_opal_resource:
            if stripped_line.startswith(COMMENT):
                logging.debug(f"Found comment line: {stripped_line}")
                new_lines.append(line)
                continue
            
            if count_bracket_changes(stripped_line) != 0:
                bracket_count += count_bracket_changes(stripped_line)

            if bracket_count == 0:  # Closing bracket of the resource block
                logging.debug(f"Found closing bracket of resource block: {stripped_line}")
                in_opal_resource = False
                if len(wrapped_block_data) == 0 and len(nested_blocks) == 0:
                    logging.info(f"Resource block does not need to be updated")
                    new_lines.append(line)
                    continue
                new_lines.append(TAB + WRAPPED_BLOCK_NAME + " {")
                new_lines.extend(wrapped_block_data)
                new_lines.extend(nested_blocks)
                new_lines.append(TAB + "}")
                new_lines.append(line)
                out = '\n'.join([str(item) for item in new_lines])
                logging.info(f"Finished updating resource block: {out}")
                wrapped_block_data = []
                nested_blocks = []
                continue

            if is_field_to_wrap(stripped_line):
                logging.debug(f"Found field to wrap: {stripped_line}")
                wrapped_block_data.append(TAB + line)
                continue

            if is_start_of_nested_block(stripped_line):
                logging.debug(f"Found start of nested block: {stripped_line}")
                nested_blocks.append(TAB + line)
                new_nested_blocks = extract_nested_block(lines, indentation_level)
                out = '\n'.join([str(item) for item in new_nested_blocks])
                logging.debug(f"Finished processing nested block: {line}\n{out}")
                nested_blocks.extend(new_nested_blocks)
                bracket_count -= 1
                continue
            
        new_lines.append(line)
    return "\n".join(new_lines)

def create_output_folder(folder_name="migration_autogen"):
    if not os.path.exists(folder_name):
        os.makedirs(folder_name)

def process_tf_files():
    create_output_folder()
    tf_files = glob.glob("*.tf")

    for tf_file in tf_files:
        logging.info(f"Processing {tf_file}")
        with open(tf_file, "r") as f:
            original_str = f.read()

        modified_str = modify_terraform_with_nested_blocks(original_str)

        output_file_path = os.path.join("migration_autogen", tf_file)
        with open(output_file_path, "w") as f:
            logging.info(f"Writing to {output_file_path}")
            f.write(modified_str)

if __name__ == "__main__":
    process_tf_files()
