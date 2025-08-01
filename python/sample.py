# sample.py

def read_file(file_path):
    try:
        with open(file_path, 'r') as file:
            content = file.read()
            return content
    except FileNotFoundError:
        return "File not found."

if __name__ == "__main__":
    file_path = "intc.txt"
    content = read_file(file_path)
    lines = content.split('\n')
    trimmed_lines = [line.strip() for line in lines]
    trimmed_lines = [line.replace(',', '') for line in trimmed_lines]
    grouped_lines = [','.join(trimmed_lines[i:i+5]) for i in range(0, len(trimmed_lines), 5)]
    content = '\n'.join(grouped_lines)

    print(content)