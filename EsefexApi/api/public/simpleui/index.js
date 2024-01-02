async function init() {
    const soundsDiv = document.getElementById('sounds');
    const userTokenInput = document.getElementById('userTokenInput');
    userTokenInput.value = document.cookie.split('; ').find(row => row.startsWith('User-Token')).split('=')[1];

    userTokenInput.addEventListener('change', (e) => {
        setUserToken(e.target.value);
    });

    let serverRequest = await fetch('/api/server', {
        method: 'GET',
        headers: {
            'User-Token': userTokenInput.value
        }
    });

    if (serverRequest.status != 200) {
        let errorP = document.createElement('p');
        errorP.innerText = 'Error: ' + serverRequest.status + ' ' + serverRequest.statusText + '\n' + await serverRequest.text();
        soundsDiv.appendChild(errorP);
        return;
    }

    let soundsRequest = await fetch(`/api/sounds/${await serverRequest.text()}`, {
        method: 'GET',
        headers: {
            'User-Token': userTokenInput.value
        }
    });
    let sounds = await soundsRequest.json();

    sounds.forEach(sound => {
        let soundButton = document.createElement('button');
        soundButton.innerText = sound.name;
        soundButton.addEventListener('click', async () => {
            await fetch(`/api/playsound/${sound.id}`, {
                method: 'POST',
                headers: {
                    'User-Token': userTokenInput.value
                }
            });
        });
        soundsDiv.appendChild(soundButton);
    });
}

function setUserToken(token) {
    document.cookie = `User-Token=${token}`;
}

init();