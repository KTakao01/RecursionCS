<!--
   問題文
   以下のエディターにHTMLで、target-containerの中に2つのdivをJavaScriptを使って作成してください。1つ目のdivにはsportというh2タグとtennis、baseballというpタグが含まれています。2つ目のdivにはanimalというh2タグとdog、catというpタグが含まれています。見本を参照してください。
-->
<div>
   <div id="target-container">

   </div>
</div>

// ここからJavaScriptを記述してください。
//div要素1,2を返す
let parentDiv1 = document.createElement("div");
let parentDiv2 = document.createElement("div");

//それぞれh2,p要素を返す
let h21 = document.createElement("h2");
let h22 = document.createElement("h2");

let p11 = document.createElement("p");
let p21 = document.createElement("p");

let p12 = document.createElement("p");
let p22 = document.createElement("p");

//h2,p要素の中身を上書きする
h21.innerHTML = "sport";
h22.innerHTML = "animal";

p11.innerHTML = "tennis";
p21.innerHTML = "baseball";

p12.innerHTML = "dog";
p22.innerHTML = "cat";

//div要素の中にh2,h２要素の中にp要素を追加する
parentDiv1.append(h21);
parentDiv2.append(h22);

parentDiv1.append(p11);
parentDiv1.append(p21);
parentDiv2.append(p12);
parentDiv2.append(p22);

//target-containerの中に子要素としてdivを追加
document.getElementById("target-container").append(parentDiv1);
document.getElementById("target-container").append(parentDiv2);