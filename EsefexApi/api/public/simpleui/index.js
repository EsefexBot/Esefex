async function init() {
    const soundsDiv = document.getElementById('sounds');
    const userTokenInput = document.getElementById('userTokenInput');

    let serverRequest = await fetch('/api/server', {
        method: 'GET',
        credentials: 'same-origin',
    });

    if (serverRequest.status != 200) {
        let errorP = document.createElement('p');
        errorP.innerText = 'Error: ' + serverRequest.status + ' ' + serverRequest.statusText + '\n' + await serverRequest.text();
        soundsDiv.appendChild(errorP);
        return;
    }

    let soundsRequest = await fetch(`/api/sounds/${await serverRequest.text()}`, {
        method: 'GET',
        credentials: 'same-origin',
    });
    let sounds = await soundsRequest.json();

    sounds.forEach(sound => {
        let soundButton = document.createElement('button');
        soundButton.innerText = sound.name;
        soundButton.addEventListener('click', async () => {
            await fetch(`/api/playsound/${sound.id}`, {
                method: 'POST',
                credentials: 'same-origin',
            });
        });
        soundsDiv.appendChild(soundButton);
    });
}

init();
