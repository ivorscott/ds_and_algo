"""palindrome module. Contains function: isPalindrome()"""

def main(): 
    print(isPalindrome("lol"))
    print(isPalindrome("scott"))

def isPalindrome(str):
    pcs = [char for char in str]
    pcs.reverse()
    return str == "".join(pcs)

if __name__ == '__main__':
    main()