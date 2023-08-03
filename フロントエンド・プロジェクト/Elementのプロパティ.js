<!--
   問題文
   以下のエディターにHTMLが書かれていますが、表記に誤りがあります。特定のIDを持つ要素を取得し、要素を正しい情報で上書きしてみましょう。
-->
<div>
   <h1>Successful startup companies in Bay Area</h1>
   <div>
      <p id="apple">Apple offers electric vehicles with clean energy</p>
      <p id="netflix">Netflix offers social platform where people can hang out</p>
      <p id="facebook">Facebooks offers the world's most advanced operating system</p>
      <p id="tesla">Tesla offers entertainment streaming services</p>
   </div>
</div>

// ここからJavaScriptを記述してください。
const apple = document.getElementById("apple").innerHTML = "Apple offers the world's most advanced operating system";
console.log(apple);

const netflix = document.getElementById("netflix").innerHTML = "Netflix offers entertainment streaming services";
console.log(netflix);

const facebook = document.getElementById("facebook").innerHTML = "Facebook offers social platform where people can hang out";
console.log(apple);

const tesla = document.getElementById("tesla").innerHTML = "Tesla offers electric vehicles with clean energy";
console.log(netflix);