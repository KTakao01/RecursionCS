// ここからJavaScriptを記述してください。
// 関数名: animalImgUrlTable
function animalImgUrlTable (animal) {
    const animalTable = {
        "dog": "https://cdn.pixabay.com/photo/2017/09/25/13/12/dog-2785074__340.jpg"
        "cat": "https://cdn.pixabay.com/photo/2014/03/29/09/17/cat-300572__340.jpg"
        "bird": "https://cdn.pixabay.com/photo/2017/02/07/16/47/kingfisher-2046453__340.jpg"
    }

    const defaultUrl = 
        "the other animals": "https://cdn.pixabay.com/photo/2014/04/05/11/20/forest-315184__340.jpg"

    return animalTable[animal] ==! undefined ? animalTable[animal] : defaultUrl;

}

function renderAnimal (animal, animalBox) {
    let div = document.createElement("div");
    let name = document.createElement("h2");
    let img = document.createElement("img");

    // 上のanimalImgUrlTable関数を完成させてください。
    // animalの名前を渡します。
    img.src = animalImgUrlTable(animal);

    name.innerHTML = "Our animal is " + animal;

    img.classList.add("animalImg");

    div.append(name);
    div.append(img);
    div.classList.add("animalDiv");

    animalBox.append(div);
}

let animalBox = document.getElementById("animal-box");

renderAnimal("dog", animalBox);
renderAnimal("cat", animalBox);
renderAnimal("bird", animalBox);
renderAnimal("lion", animalBox);
renderAnimal("bear", animalBox);
renderAnimal("fox", animalBox);