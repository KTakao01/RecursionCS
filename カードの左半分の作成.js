// <div class="vh-100 bg-dark d-flex flex-column justify-content-center align-items-center">
//     <div class="d-flex align-items-center col-md-7 col-10 m-1">
//         <div class="d-flex col-12 profile-card">
//
//             左半分
//             <div class="col-8 py-3">
//                 <h4>Kaiden Herman</h4>
//                 <div class="py-2">
//                     <p>Job : <br>Software Engineer</p>
//                 </div>
//                 <div class="py-2">
//                     <p>Skill : <br>C++, C#, Java, PHP, JavaScript, Python</p>
//                 </div>
//                 <div class="py-2">
//                     <p>Country :<br>United States</p>
//                 </div>
//             </div>

//         </div>
//     </div>
// </div>

let innerFlex = document.createElement("div");
innerFlex.classList.add("d-flex", "align-items-center", "col-md-7", "col-10", "m-1");

let cardDiv = document.createElement("div");
innerFlex.append(cardDiv);

cardDiv.classList.add("d-flex", "col-12", "profile-card");

// 今回、py-2が入った要素が3つ出てくるのでdiv1のコピーを作成します。
// 要素のノード.cloneNode(true)はノードの完全コピーを生成します。
// div2 = div1 というのはオブジェクト参照をコピーすることを意味するので注意してください。

// ここからJavaScriptを記述してください。
// 最後にコンソールに出力して確認しましょう。
let leftInfo = document.createElement("div");
leftInfo.classList.add("col-8","py-3");

let nameTitle = document.createElement("h4");
nameTitle.innerHTML = "Kaiden Herman";

let div1 = document.createElement("div1");
div1.classList.add("py-2");

div2 = div1.cloneNode(true);
div3 = div1.cloneNode(true);

let employeeJob = document.createElement("p");
let employeeSkill = document.createElement("p");
let employeeCountry = document.createElement("p");

employeeJob.innerHTML = "Job : <br>Software Engineer"
employeeSkill.innerHTML = "Skill : <br>C++, C#, Java, PHP, JavaScript, Python"
employeeCountry.innerHTML = "Country :<br>United States"

cardDiv.append(leftInfo);

leftInfo.append(nameTitle);
leftInfo.append(div1);
leftInfo.append(div2);
leftInfo.append(div3);

div1.append(employeeJob);
div2.append(employeeSkill);
div3.append(employeeCountry);

console.log(leftInfo);


