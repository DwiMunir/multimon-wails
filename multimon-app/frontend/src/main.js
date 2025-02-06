import './style.css';
import { Greet, ShowMessage } from '../wailsjs/go/main/App';

// document.querySelector('#app').innerHTML = `
//   <div class="container">
//     <div class="content">
//       <div class="webview-container">
//         <iframe 
//           src="https://cxds-player.mdevoffice.net" 
//           id="webview" 
//           frameborder="0" 
//           sandbox="allow-scripts allow-same-origin allow-forms"
//           allowfullscreen>
//         </iframe>
//       </div>
//     </div>
//   </div>
// `

document.location.href = 'https://cxds-player.mdevoffice.net';

// Setup event listeners
document.getElementById('greet').addEventListener('click', async () => {
    const nameInput = document.getElementById('name');
    const result = document.getElementById('result');
    
    try {
        const greeting = await Greet(nameInput.value);
        result.textContent = greeting;
    } catch (err) {
        console.error(err);
    }
});

// Listen for messages from the backend
window.runtime.EventsOn("show-message", (message) => {
    const result = document.getElementById('result');
    result.textContent = message;
});

let isDragging = false;
let offsetX = 0;
let offsetY = 0;

const dragArea = document.getElementsByClassName("dragArea");

dragArea.addEventListener("mousedown", (event) => {
  isDragging = true;
  offsetX = event.clientX;
  offsetY = event.clientY;
});

document.addEventListener("mousemove", (event) => {
  if (isDragging) {
    const deltaX = event.clientX - offsetX;
    const deltaY = event.clientY - offsetY;

    // Memindahkan jendela dengan moveBy
    window.moveBy(deltaX, deltaY);

    offsetX = event.clientX;
    offsetY = event.clientY;
  }
});

document.addEventListener("mouseup", () => {
  isDragging = false;
});
