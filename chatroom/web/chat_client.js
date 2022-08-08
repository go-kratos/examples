const messages = document.getElementById("messages");
const form = document.getElementById("msgForm");
const inputBox = document.getElementById("inputBox");

function showMessage(message) {
    const msgDiv = document.createElement("div");
    msgDiv.classList.add("msgCtn");

    msgDiv.innerHTML = message;
    messages.appendChild(msgDiv);
}

form.addEventListener("submit", (event) => {
    event.preventDefault();
    const message = inputBox.value;
    console.log('send message: ' + message);

    sendChatMessage(message);

    inputBox.value = "";
});

function sendChatMessage(message) {
    let packet = {
        message: message,
        sender: "",
        timestamp: "",
    };

    const str = JSON.stringify(packet);
    const buff = new TextEncoder().encode(str);

    sendMessage(0, buff);
}

function sendMessage(id, payload) {
    let buff = new Uint8Array(4 + payload.byteLength);
    let dv = new DataView(buff.buffer);
    dv.setInt32(0, id);
    buff.set(payload, 4);

    console.log(ab2str(buff))

    ws.send(dv.buffer);
}

function ab2str(buffer) {
    const uint8 = new Uint8Array(buffer);
    const decoder = new TextDecoder('utf8');
    return decoder.decode(uint8)
}

function str2ab(str) {
    const buf = new ArrayBuffer(str.length * 2); // 2 bytes for each char
    const bufView = new Uint16Array(buf);
    let i = 0, strLen = str.length;
    for (; i < strLen; i++) {
        bufView[i] = str.charCodeAt(i);
    }
    return buf;
}

function handleMessage(messageType, messagePayload) {
    console.log('Type', messageType);
    console.log('Payload:', messagePayload);

    const str = ab2str(messagePayload);
    console.log('Message:', str);

    let packet = JSON.parse(str);
    console.log('recv:', packet);
    showMessage(packet.message);
}

let ws;

function createWebsocketClient() {
    if (ws) {
        ws.onerror = ws.onopen = ws.onclose = null;
        ws.close();
    }

    ws = new WebSocket(`ws://localhost:7700`);
    ws.binaryType = "arraybuffer";

    ws.onerror = function () {
        showMessage('WebSocket error');
    };

    ws.onopen = () => {
        showMessage('WebSocket connection established')
        console.log('WebSocket connection established')
    }

    ws.onclose = () => {
        showMessage('WebSocket connection closed')
        console.log('WebSocket connection closed')
        ws = null;
    }

    ws.addEventListener("ping", () => {
        console.log("ping");
    });
    ws.addEventListener("pong", () => {
        console.log("pong");
    });

    ws.onmessage = function (event) {
        const dv = new DataView(event.data);
        const messageType = dv.getInt32(0);
        handleMessage(messageType, event.data.slice(4));
    };
}

createWebsocketClient();
