
<!-- 
    問題文
    動物を表す文字列と入れたい箇所を受け取って、画像と説明を追加するrenderAnimalという関数を作成してください。

    url
    dog: https://cdn.pixabay.com/photo/2017/09/25/13/12/dog-2785074__340.jpg
    cat: https://cdn.pixabay.com/photo/2014/03/29/09/17/cat-300572__340.jpg
    bird: https://cdn.pixabay.com/photo/2017/02/07/16/47/kingfisher-2046453__340.jpg
    the other animals: https://cdn.pixabay.com/photo/2014/04/05/11/20/forest-315184__340.jpg
-->
<head>
    <!-- Bootstrapの読み込み -->
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.5.0/css/bootstrap.min.css" integrity="sha384-9aIt2nRpC12Uk9gS9baDl411NQApFmC26EwAOH8WgZl5MYYxFfc+NcPb1dKGj7Sk" crossorigin="anonymous">
</head>
<div>
    <div class="row" id="animal-box">

    </div>
</div>

.animalImg{
    width: 100px;
    height: 100px;
    padding: 1em;
}

.animalDiv{
    display: flex;
    align-content: center;
    justify-content: center;
    width: 50%;
    margin-top:2em;
}

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