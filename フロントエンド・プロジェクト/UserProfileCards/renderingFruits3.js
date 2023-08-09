function animalImgUrl (animal) {
    animal = animal.toLowerCase();

    if (animal == "dog") {
        return "https://cdn.pixabay.com/photo/2017/09/25/13/12/dog-2785074__340.jpg";
    } else if (animal == "cat") {
        return "https://cdn.pixabay.com/photo/2014/03/29/09/17/cat-300572__340.jpg";
    } else if (animal == "bird") {
        return "https://cdn.pixabay.com/photo/2017/02/07/16/47/kingfisher-2046453__340.jpg";
    } else {
        return "https://cdn.pixabay.com/photo/2014/04/05/11/20/forest-315184__340.jpg";
    }
}

// ここからJavaScriptを記述してください。
// 関数名: renderAnimal
// 見本のHTMLを参考にしてください。
function renderAnimal (animal,ele){
    //pタグを作成する
    let animalP = document.createElement("p");
    animalP.innerHTML = "Our animal is " + animal; 

    //imgタグを作成する
    let animalImg = document.createElement("img")
    animalImg.src = animalImgUrl(animal);

    //imgタグにクラスを追加
    animalImg.classList("animalImg")

    //親要素としてdivタグを作成する
    let animalDiv = document.createElement("div");
    //divタグにクラスを追加
    animalDiv.classList("animalDiv")
    //親要素にp,imgを組み入れる
    animalDiv.append(animalP);
    animalDiv.append(animalImg);

    //上記要素を子要素として元々のdivに組み入れる
    ele.append(animalDiv);
}

let animalBox = document.getElementById("animal-box")

renderAnimal(dog,animalBox);
renderAnimal(cat,animalBox);
renderAnimal(bird,animalBox);
renderAnimal(lion,animalBox);
renderAnimal(bear,animalBox);
renderAnimal(fox,animalBox);