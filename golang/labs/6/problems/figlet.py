from pyfiglet import Figlet
import sys

figlet = Figlet()

fontList = figlet.getFonts()
s = input("input: ")
if len(sys.argv) < 2:
    figlet.setFont(font = "briteb")
    print(figlet.renderText(s))
elif (len(sys.argv) == 2) & (sys.argv[1] in fontList):
    figlet.setFont(font = sys.argv[1])
    print(figlet.renderText(s))
else:
    print("wrong argument!")
    sys.exit(1)

