import re

def replace_roadmap():
    with open('ROADMAP.md', 'r') as f:
        content = f.read()

    # Replacements
    replacements = [
        (r'\bpresets\b', 'items'),
        (r'\bPresets\b', 'Items'),
        (r'\bpreset\b', 'item'),
        (r'\bPreset\b', 'Item'),
        (r'AI guidance items for agentic', 'AI guidance items for agentic'),
        (r'sheave.presets', 'sheave.items'),
    ]

    for pattern, repl in replacements:
        content = re.sub(pattern, repl, content)

    # Some manual fixes if needed
    content = content.replace("sheave.items module", "sheave.items module")
    
    with open('ROADMAP.md', 'w') as f:
        f.write(content)

if __name__ == '__main__':
    replace_roadmap()
