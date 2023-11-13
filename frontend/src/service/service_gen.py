import sys

name = sys.argv[1]
tmpl_file = sys.argv[2]
keyword = sys.argv[3]

camel_name = name[0].lower() + name[1:]
pascal_name = name[0].upper() + name[1:]
camel_key = keyword[0].lower() + keyword[1:]
pascal_key = keyword[0].upper() + keyword[1:]

with open(tmpl_file, 'r') as f:
    tmpl = f.readlines()
    with open(pascal_name + 'Service.ts', 'w') as w:
        for line in tmpl:
            line = line.replace(camel_key, camel_name)
            line = line.replace(pascal_key, pascal_name)
            w.write(line)