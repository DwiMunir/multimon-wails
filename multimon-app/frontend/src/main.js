import './style.css';
import { Greet, ShowMessage } from '../wailsjs/go/main/App';

document.querySelector('#app').innerHTML = `
  <div class="container">
    <div class="content">
      <div class="webview-container">
        <iframe 
          src="https://cxds-player.mdevoffice.net" 
          id="webview" 
          frameborder="0" 
          allowfullscreen>
        </iframe>
      </div>
    </div>
  </div>
`

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
