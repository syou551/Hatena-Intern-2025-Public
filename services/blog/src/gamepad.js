window.addEventListener("gamepadconnected", (e) => {
  // ゲームパッドが接続された時に，「接続されました」とポップアップ要素を表示する
  const popup = document.createElement("div");
  popup.textContent = "ゲームパッドが接続されました";
  popup.style.position = "fixed";
  popup.style.top = "50%";
  popup.style.left = "50%";
  popup.style.transform = "translate(-50%, -50%)";
  popup.style.backgroundColor = "rgba(0, 0, 0, 0.8)";
  popup.style.color = "white";
  popup.style.padding = "10px 20px";
  popup.style.borderRadius = "5px";
  popup.style.zIndex = "1000"; // 他の要素の上に表示
  popup.style.fontSize = "20px";
  // 閉じるボタンを追加
  const closeButton = document.createElement("button");
  closeButton.textContent = "閉じる";
  closeButton.style.marginLeft = "10px";
  closeButton.style.backgroundColor = "red";
  closeButton.style.color = "white";
  closeButton.style.border = "none";
  closeButton.style.padding = "5px 10px";
  closeButton.style.cursor = "pointer";
  closeButton.addEventListener("click", () => {
    popup.remove();
    showCracker();
  });

  popup.appendChild(closeButton);
  document.body.appendChild(popup);
  
});

function showCracker() {
  // index.htmlに全画面でクラッカーを表示する
  // クラッカーは/imagesから取得して，img要素で表示する
  const img = document.createElement("img");
  img.src = "/images/cracker.png";
  img.style.position = "fixed";
  img.style.top = "50%";
  img.style.left = "50%";
  img.style.transform = "translate(-50%, -50%)";
  img.style.zIndex = "1000"; 
  img.style.width = "300px"; 
  img.style.height = "300px";
  //　背景は黒色で透過にする
  img.style.backgroundColor = "rgba(0, 0, 0, 0.5)";
  img.style.boxShadow = "0 0 10px rgba(0, 0, 0, 0.5)"; // 影をつける

  document.body.appendChild(img);
  // 音声を再生する
  const audioElement = document.querySelector("audio");
  audioElement.currentTime = 0;
  audioElement.play();

  // 5秒後に削除する
  setTimeout(() => {
    img.remove();
  }, 5000);
}