document.addEventListener('DOMContentLoaded', async function() {
    await loadDevices();
});

document.getElementById('create-device-form').addEventListener('submit', async function(event) {
    event.preventDefault();
    const deviceId = document.getElementById('device_id').value;

    const response = await fetch('/device', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ device_id: deviceId })
    });

    const data = await response.json();
    const responseElement = document.getElementById('create-device-response');
    const form = event.target;

    if (response.ok) {
        responseElement.textContent = 'Device created successfully!';
        responseElement.classList.add('alert', 'alert-success');
        await loadDevices();
    } else {
        responseElement.textContent = 'Failed to create new device.';
        responseElement.classList.add('alert', 'alert-danger');
    }

    form.reset();

    setTimeout(() => {
        responseElement.textContent = '';
        responseElement.classList.remove('alert', 'alert-success', 'alert-danger');
    }, 5000);
});

document.getElementById('upload-update-form').addEventListener('submit', async function(event) {
    event.preventDefault();
    const formData = new FormData(this);

    const response = await fetch('/upload_update', {
        method: 'POST',
        body: formData
    });

    const data = await response.json();
    const responseElement = document.getElementById('upload-update-response');
    const form = event.target;

    if (response.ok) {
        responseElement.textContent = 'Update uploaded successfully!';
        responseElement.classList.add('alert', 'alert-success');
    } else {
        responseElement.textContent = 'Failed to upload update.';
        responseElement.classList.add('alert', 'alert-danger');
    }

    form.reset();

    setTimeout(() => {
        responseElement.textContent = '';
        responseElement.classList.remove('alert', 'alert-success', 'alert-danger');
    }, 5000);
});

document.getElementById('get-updates-form').addEventListener('submit', function(event) {
    event.preventDefault();

    var tableBody = document.getElementById("updates-list");
    tableBody.innerHTML = "";

    var startDate = document.getElementById("start_date").value;
    var endDate = document.getElementById("end_date").value;

    var requestData = {
        start_date: startDate,
        end_date: endDate
    };

    fetch("/updates", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(requestData)
    })
    .then(response => response.json())
    .then(data => {
        for (var i = data.length - 1; i >= 0; i--) {
            var update = data[i];
            var row = tableBody.insertRow();
            var iconCell = row.insertCell(0);
            var versionCell = row.insertCell(1);
            var descriptionCell = row.insertCell(2);
            var dateCell = row.insertCell(3);

            var trashIcon = document.createElement('i');
            trashIcon.className = 'fas fa-trash-alt';
            trashIcon.addEventListener('click', function() {
                deleteUpdate(update.version);
            });
            iconCell.appendChild(trashIcon);

            versionCell.innerHTML = update.version;
            descriptionCell.innerHTML = update.description;

            // Format the time
            var date = new Date(update.date);
            var formattedDate = `${date.getFullYear()}-${(date.getMonth() + 1).toString().padStart(2, '0')}-${date.getDate().toString().padStart(2, '0')} ${date.getHours().toString().padStart(2, '0')}:${date.getMinutes().toString().padStart(2, '0')}:${date.getSeconds().toString().padStart(2, '0')}`;
            
            dateCell.innerHTML = formattedDate;
        }
    })
    .catch(error => console.error('Error:', error));
});

async function loadDevices() {
    const response = await fetch('/devices');
    const devices = await response.json();
    const devicesList = document.getElementById('devices-list');
    devicesList.innerHTML = '';

    // Reverse the order of devices array
    devices.reverse();

    devices.forEach(device => {
        const row = document.createElement('tr');

        const iconCell = document.createElement('td');
        const trashIcon = document.createElement('i');
        trashIcon.className = 'fas fa-trash-alt';
        trashIcon.addEventListener('click', function() {
            deleteDevice(device.device_id);
        });
        iconCell.appendChild(trashIcon);
        row.appendChild(iconCell);

        const deviceIdCell = document.createElement('td');
        deviceIdCell.textContent = device.device_id;
        row.appendChild(deviceIdCell);

        const deviceVersionCell = document.createElement('td');
        deviceVersionCell.textContent = device.device_version || 'N/A';
        row.appendChild(deviceVersionCell);

        const lastUpdateTimeCell = document.createElement('td');
        const lastUpdateDate = new Date(device.last_update);
        const formattedLastUpdate = lastUpdateDate.toLocaleString('en-US', {
            year: 'numeric',
            month: '2-digit',
            day: '2-digit',
            hour: '2-digit',
            minute: '2-digit',
            second: '2-digit',
        });
        lastUpdateTimeCell.textContent = formattedLastUpdate;
        row.appendChild(lastUpdateTimeCell);

        devicesList.appendChild(row);
    });
}

async function deleteDevice(deviceId) {
    const response = await fetch('/device_delete', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ device_id: deviceId })
    });

    if (response.ok) {
        await loadDevices();
    } else {
        alert('Failed to delete device.');
    }
}

async function deleteUpdate(version) {
    event.preventDefault();

    const response = await fetch('/update_delete', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ version: version })
    });

    if (response.ok) {
        // Reload updates
        document.getElementById('get-updates-form').dispatchEvent(new Event('submit'));
    } else {
        alert('Failed to delete update.');
    }
}