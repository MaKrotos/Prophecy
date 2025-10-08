/**
 * UI компоненты и сообщения для не-Telegram окружения
 */

export function showTelegramOnlyMessage(options = {}) {
  const {
    title = "Telegram App Required",
    message = "This application can only be used within the Telegram messenger.",
    buttonText = "Open in Telegram",
    telegramLink = null
  } = options;
  
  document.body.innerHTML = `
    <div style="
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
      height: 100vh;
      text-align: center;
      padding: 20px;
      font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
      background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
      color: white;
    ">
      <div style="background: rgba(255,255,255,0.1); padding: 40px; border-radius: 20px; backdrop-filter: blur(10px);">
        <h1 style="margin-bottom: 20px; font-size: 24px;">${title}</h1>
        <p style="margin-bottom: 30px; font-size: 16px; opacity: 0.9;">${message}</p>
        ${telegramLink ? `
          <a href="${telegramLink}" style="
            display: inline-block;
            padding: 12px 24px;
            background: #0088cc;
            color: white;
            text-decoration: none;
            border-radius: 10px;
            font-weight: 500;
            transition: background 0.3s;
          " onmouseover="this.style.background='#0077b3'" onmouseout="this.style.background='#0088cc'">
            ${buttonText}
          </a>
        ` : ''}
      </div>
    </div>
  `;
}