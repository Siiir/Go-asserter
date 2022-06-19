import re, os

forbiddenFiles= ["cmd.exe","go.sum","go.mod",".git",".gitattributes","LICENSE","lineCounter.py","nameChanger.py"]
patt= re.compile(r"\bNilFail\b")
for fName in os.listdir():
    if fName in forbiddenFiles: continue
    
    try:
        fi= open(fName, encoding= "UTF-8")
        s= fi.read()
        s= patt.sub("NilFailError",s)
        with open(fName,"wt", encoding= "UTF-8") as fi:
            fi.write(s)
    except UnicodeDecodeError:
        print(f"Caught `UnicodeDecodeError` while dealing with {fName}.")