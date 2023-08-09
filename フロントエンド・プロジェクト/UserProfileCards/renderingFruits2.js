
// ここからJavaScriptを記述してください。
// 関数名: animalImgUrl
function animalImgUrl(animal){
    animal = animal.toLowerCase();

    if(animal === "dog")return "https://cdn.pixabay.com/photo/2017/09/25/13/12/dog-2785074__340.jpg";
    else if(animal === "cat")return "https://cdn.pixabay.com/photo/2014/03/29/09/17/cat-300572__340.jpg";
    else if(animal === "bird")return "https://cdn.pixabay.com/photo/2017/02/07/16/47/kingfisher-2046453__340.jpg";
    else return "https://cdn.pixabay.com/photo/2014/04/05/11/20/forest-315184__340.jpg";

}

//divタグを作成して、
let animalDiv = document.createElement("div");

//divタグの中にpとimgタグを宣言する
let animalP = document.createElement("p");
let animalImg = document.createElement("img");


//pとimgの中身を決める
animalP.innerHTML = "Our animal is dog";
animalImg.src = animalImgUrl("dog");

//CSSを効かせるためにクラスを設定
animalImg.classList.add("animalImg");
animalDiv.classList.add("animalDiv");


//divタグの中にpとimgタグをいれる
animalDiv.append(animalP);
animalDiv.append(animalImg);

//子要素としてdivをいれる
document.getElementById("animal-container").append(animalDiv);

