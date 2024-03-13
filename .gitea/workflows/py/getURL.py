import os

tag = os.getenv('GITHUB_SERVER_URL')
if tag:
    tag=tag.replace('https://', '')
    tag = "URL=" + tag
    print(tag)
