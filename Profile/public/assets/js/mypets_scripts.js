document.addEventListener('DOMContentLoaded', function () {
    // �������� DOM
    const petsList = document.getElementById('petsList');
    const addPetBtn = document.querySelector('.add-pet-btn');

    // ���������, ���� �� ������� � localStorage ��� ��������
    loadPetsFromStorage();

    // ���������� ������ ���������� �������
    addPetBtn.addEventListener('click', function (e) {
        e.preventDefault(); // ������������� ������� �� ������
        showAddPetModal();
    });

    // ������� �������� �������� �� localStorage
    function loadPetsFromStorage() {
        const savedPets = JSON.parse(localStorage.getItem('pets')) || [];

        if (savedPets.length > 0) {
            // ������� ��������� � ������ ������, ���� ��� ����
            const emptyMessage = petsList.querySelector('.empty-pets-message');
            if (emptyMessage) petsList.removeChild(emptyMessage);

            // ��������� ������� ������� �� ���������
            savedPets.forEach(pet => {
                addPetToDOM(pet);
            });
        }
    }

    // ������� ���������� ������� � DOM
    function addPetToDOM(pet) {
        // ������� ��������� � ������ ������, ���� ��� ����
        const emptyMessage = petsList.querySelector('.empty-pets-message');
        if (emptyMessage) petsList.removeChild(emptyMessage);

        // ������� �������� �������
        const petCard = document.createElement('div');
        petCard.className = 'pet-card';
        petCard.dataset.id = pet.id;
        petCard.onclick = function () { window.location.href = '/pet/' + pet.id; };

        petCard.innerHTML = `
            <div class="pet-icon-circle">
                <img src="${pet.imageUrl || './assets/img/free-icon-pets-12452249.png'}" alt="${pet.type}">
            </div>
            <div class="pet-info">
                <p><strong>������:</strong> ${pet.name}</p>
                <p><strong>���:</strong> ${pet.type}</p>
                <p><strong>������:</strong> ${pet.breed}</p>
                <p><strong>���:</strong> ${pet.chip}</p>
                <p><strong>�������:</strong> ${pet.age}</p>
            </div>
            <button class="pet-profile-btn">�������</button>
            <button class="delete-pet-btn">�������</button>
        `;

        // ��������� ���������� ��������
        const deleteBtn = petCard.querySelector('.delete-pet-btn');
        deleteBtn.addEventListener('click', function (e) {
            e.stopPropagation(); // ������������� ������������ ����� �� ��������
            deletePet(pet.id);
        });

        petsList.appendChild(petCard);
    }

    // ������� �������� �������
    function deletePet(petId) {
        if (confirm('�� �������, ��� ������ ������� ����� �������?')) {
            // ������� �� DOM
            const petCard = document.querySelector(`.pet-card[data-id="${petId}"]`);
            if (petCard) petCard.remove();

            // ������� �� localStorage
            const savedPets = JSON.parse(localStorage.getItem('pets')) || [];
            const updatedPets = savedPets.filter(pet => pet.id !== petId);
            localStorage.setItem('pets', JSON.stringify(updatedPets));

            // ���� �������� �� ��������, ���������� ���������
            if (updatedPets.length === 0) {
                showEmptyMessage();
            }
        }
    }

    // ������� ������ ���������� ���� ���������� �������
    function showAddPetModal() {
        // ������� ��������� ����
        const modal = document.createElement('div');
        modal.className = 'modal';
        modal.innerHTML = `
            <div class="modal-content">
                <span class="close-modal">&times;</span>
                <h2>�������� ������ �������</h2>
                <form id="addPetForm">
                    <div class="form-group">
                        <label for="petName">������:</label>
                        <input type="text" id="petName" required>
                    </div>
                    <div class="form-group">
                        <label for="petType">���:</label>
                        <select id="petType" required>
                            <option value="�����">�����</option>
                            <option value="������">������</option>
                        </select>
                    </div>
                    <div class="form-group">
                        <label for="petBreed">������:</label>
                        <input type="text" id="petBreed" required>
                    </div>
                    <div class="form-group">
                        <label for="petChip">����� ����:</label>
                        <input type="text" id="petChip">
                    </div>
                    <div class="form-group">
                        <label for="petAge">�������:</label>
                        <input type="text" id="petAge" required>
                    </div>
                    <div class="form-group">
                        <label for="petImage">������ �� ���� (�������������):</label>
                        <input type="text" id="petImage">
                    </div>
                    <button type="submit" class="submit-btn">�������� �������</button>
                </form>
            </div>
        `;

        document.body.appendChild(modal);

        // ���������� �������� ���������� ����
        const closeBtn = modal.querySelector('.close-modal');
        closeBtn.addEventListener('click', function () {
            document.body.removeChild(modal);
        });

        // ���������� �������� �����
        const form = modal.querySelector('#addPetForm');
        form.addEventListener('submit', function (e) {
            e.preventDefault();
            handleAddPetForm();
        });
    }

    // ���������� ����� ���������� �������
    function handleAddPetForm() {
        // �������� ������ �� �����
        const petData = {
            id: Date.now().toString(),
            name: document.getElementById('petName').value,
            type: document.getElementById('petType').value,
            breed: document.getElementById('petBreed').value,
            chip: document.getElementById('petChip').value,
            age: document.getElementById('petAge').value,
            imageUrl: document.getElementById('petImage').value || null
        };

        // ��������� �������
        addNewPet(petData);

        // ��������� ��������� ����
        const modal = document.querySelector('.modal');
        if (modal) document.body.removeChild(modal);
    }

    // ������� ���������� ������ �������
    function addNewPet(petData) {
        // ��������� � DOM
        addPetToDOM(petData);

        // ��������� � localStorage
        const savedPets = JSON.parse(localStorage.getItem('pets')) || [];
        savedPets.push(petData);
        localStorage.setItem('pets', JSON.stringify(savedPets));
    }

  
});
