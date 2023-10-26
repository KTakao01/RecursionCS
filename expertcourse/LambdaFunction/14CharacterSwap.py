##https://recursionist.io/dashboard/problems/439
##文字で構成される配列 charList が与えられるので、大文字は小文字に、小文字は大文字に変換し新しい配列を返す swapCase という関数を map 関数を使用して作成し、テストケースを出力してください。
def swapCase(charList):
    # 関数を完成させてください
    return list(map(lambda x:x.lower() if 'A' <= x and 'Z' >= x else x.upper()  , charList))