//自分の回答
function fruitImgUrl(fruit){
    // 関数を完成させてください
    switch(fruit){
        case "banana" :
            return "url1";
        case "BANANA" :
            return "url1";
        case "pineapple" :
            return "url2";
        case "PINEAPPLE" :
            return "url2"

        case "pear":
            return "url3";
        case "PEAR":
            return "url3";

        default:
            return "no image";
    }
}

//回答例
const fruits = {
    banana: 'url1',
    pineapple: 'url2',
    pear: 'url3'
}

function fruitImgUrl(fruit) {
    const lowerCaseFruit = fruit.toLowerCase()
    const url = fruits[lowerCaseFruit]

    return url || 'no image'
}