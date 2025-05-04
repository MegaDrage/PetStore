// js/inventory.js

let medicines = JSON.parse(localStorage.getItem('medicines')) || [
    'Амоксициллин',
    'Цефтриаксон',
    'Ибупрофен',
    'Парацетамол',
    'Лоперамид',
    'Омепразол'
];

let autoOrderEnabled = false;

document.addEventListener('DOMContentLoaded', function() {
    initApp();
});

function initApp() {
    document.querySelector('.auto-order-btn')?.addEventListener('click', function() {
        autoOrderEnabled = !autoOrderEnabled;
        this.textContent = autoOrderEnabled ? 'Выключить автозаказ' : 'Включить автозаказ';
        if (autoOrderEnabled) autoOrderMedicines();
    });

    document.getElementById('medicine-search')?.addEventListener('input', function(e) {
        updateAutocomplete(e.target);
    });

    document.querySelectorAll('.order-btn').forEach(btn => {
        btn.addEventListener('click', handleOrderClick);
    });
}

function updateAutocomplete(input) {
    const suggestions = document.querySelector('.autocomplete-suggestions');
    suggestions.innerHTML = '';
    const inputVal = input.value.toLowerCase();
    
    if (!inputVal) {
        suggestions.style.display = 'none';
        return;
    }

    const filtered = medicines
        .filter(item => item.toLowerCase().includes(inputVal))
        .slice(0, 5);

    if (filtered.length) {
        filtered.forEach(item => {
            const div = document.createElement('div');
            div.className = 'autocomplete-item';
            div.textContent = item;
            div.onclick = () => {
                input.value = item;
                suggestions.style.display = 'none';
            };
            suggestions.appendChild(div);
        });
        suggestions.style.display = 'block';
    } else {
        suggestions.style.display = 'none';
    }
}

function autoOrderMedicines() {
    document.querySelectorAll('.inventory-table tbody tr').forEach(row => {
        const statusCell = row.querySelector('td:nth-child(4)');
        if (statusCell.classList.contains('warning')) {
            const quantityCell = row.querySelector('td:nth-child(2)');
            const minStock = parseInt(row.querySelector('td:nth-child(3)').textContent);
            let quantity = parseInt(quantityCell.textContent) + 20;
            quantityCell.textContent = quantity;
            if (quantity >= minStock) {
                statusCell.className = 'ok';
                statusCell.textContent = 'В норме';
            } else {
                statusCell.className = 'warning';
                statusCell.textContent = 'Недостаточно';
            }
            const orderBtn = row.querySelector('.order-btn');
            orderBtn.disabled = quantity >= minStock;
            const medicineName = row.querySelector('td:first-child').textContent;
            addNewMedicine(medicineName);
        }
    });
    
    if (autoOrderEnabled) {
        setTimeout(autoOrderMedicines, 5000);
    }
}

function addNewMedicine(name) {
    if (!medicines.includes(name)) {
        medicines.push(name);
        localStorage.setItem('medicines', JSON.stringify(medicines));
    }
}

function handleOrderClick() {
    const row = this.closest('tr');
    const medicineName = row.querySelector('td:first-child').textContent;
    const quantityCell = row.querySelector('td:nth-child(2)');
    let quantity = parseInt(quantityCell.textContent) + 10;
    quantityCell.textContent = quantity;
    const minStock = parseInt(row.querySelector('td:nth-child(3)').textContent);
    const statusCell = row.querySelector('td:nth-child(4)');
    if (quantity >= minStock) {
        statusCell.className = 'ok';
        statusCell.textContent = 'В норме';
    } else {
        statusCell.className = 'warning';
        statusCell.textContent = 'Недостаточно';
    }
    this.disabled = quantity >= minStock;
    addNewMedicine(medicineName);
    alert(`Заказано 10 единиц: ${medicineName}`);
}

function searchMedicine() {
    const input = document.getElementById('medicine-search');
    const searchTerm = input.value.toLowerCase();
    document.querySelectorAll('.inventory-table tbody tr').forEach(row => {
        const name = row.querySelector('td:first-child').textContent.toLowerCase();
        if (name.includes(searchTerm)) {
            row.style.display = '';
        } else {
            row.style.display = 'none';
        }
    });
}