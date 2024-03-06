function showOptions() {
    document.getElementById("options").style.display = "block";
}

function showActions(option) {
    document.getElementById("actions").style.display = "block";
    // Here you can handle different actions based on the chosen option
}

function showActionForm(action) {
    document.getElementById("actionForm").style.display = "block";
    // Here you can show the appropriate form fields based on the chosen action
}

function showMessage(message) {
    const messageElement = document.getElementById("message");
    messageElement.textContent = message;
    messageElement.style.display = "block";
    // You can hide the message after a few seconds if needed
    setTimeout(() => {
        messageElement.style.display = "none";
    }, 3000);
}
