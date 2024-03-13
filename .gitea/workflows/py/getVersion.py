import os

tag = os.getenv('GITHUB_REF')
if tag:
    tag=tag.replace('refs/tags/', '')
    tag = "VERSION=" + tag
    print(tag)
