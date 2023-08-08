let innerFlex = document.createElement("div");
innerFlex.classList.add("d-flex", "align-items-center", "col-md-7", "col-10", "m-1");

let cardDiv = document.createElement("div");
innerFlex.append(cardDiv);
cardDiv.classList.add("d-flex", "col-12", "profile-card");

// 左半分
let leftInfo = document.createElement("div");
leftInfo.classList.add("col-8", "py-3");

let div1 = document.createElement("div");
div1.classList.add("py-2")

let div2 = div1.cloneNode(true);
let div3 = div1.cloneNode(true);

let nameTitle = document.createElement("h4");
nameTitle.innerHTML = "Kaiden Herman";

let employeeJob = document.createElement("p")
let employeeSkill = document.createElement("p")
let employeeCountry = document.createElement("p");

employeeJob.innerHTML = "Job: " + "<br>"  + "Software Engineer";
div1.append(employeeJob);

employeeSkill.innerHTML = "Skill: " + "<br>"  + "C++, C#, Java, PHP, JavaScript, Python";
div2.append(employeeSkill);

employeeCountry.innerHTML = "Country : " + "<br>"  + "United States";
div3.append(employeeCountry);

leftInfo.append(nameTitle);
leftInfo.append(div1);
leftInfo.append(div2);
leftInfo.append(div3);
// *** 左半分終了 ***

// 右半分
// <div class="vh-100 bg-dark d-flex flex-column justify-content-center align-items-center">
//     <div class="d-flex align-items-center col-md-7 col-10 m-1">
//         <div class="d-flex col-12 profile-card">
//
//             左半分
//             <div class="col-8 py-3">
//                 <h4>Kaiden Herman</h4>
//                 <div class="py-2">
//                     <p>Job :</p>
//                     <p>Software Engineer</p>
//                 </div>
//                 <div class="py-2">
//                     <p>Skill :</p>
//                     <p>C++, C#, Java, PHP, JavaScript, Python</p>
//                 </div>
//                 <div class="py-2">
//                     <p>Country :</p>
//                     <p>United States</p>
//                 </div>
//             </div>
//
//             右半分
//             <div class="col-4 d-flex justify-content-center align-items-center">
//                 <div>
//                     <img class="avatar" src="https://pbs.twimg.com/profile_images/501759258665299968/3799Ffxy.jpeg">
//                 </div>
//             </div>
//         </div>
//     </div>
// </div>

// ここからJavaScriptを記述してください。
// 作成したimgタグにはavatarというクラスを追加してください。
// 右要素を作成した後は、左要素と共にprofilesに追加してください。


//profile card
cardDiv.append(leftInfo);


//右半分
let rightInfo = document.createElement("div");


rightInfo.classList.add("col-4","d-flex","justify-content-center","align-items-center");
cardDiv.append(rightInfo);


let div4 = document.createElement("div");
let avatarImg = document.createElement("img");
avatarImg.classList.add("avatar");
avatarImg.src = "https://pbs.twimg.com/profile_images/501759258665299968/3799Ffxy.jpeg";

rightInfo.append(div4);
div4.append(avatarImg);

document.getElementById("profiles").append(innerFlex);
console.log(profiles);