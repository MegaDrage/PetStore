document.addEventListener('DOMContentLoaded', function () {
    // Элементы DOM
    const petsList = document.getElementById('petsList');
    const addPetBtn = document.querySelector('.add-pet-btn');

    // Проверяем, есть ли питомцы в localStorage при загрузке
    loadPetsFromStorage();

    // Обработчик кнопки добавления питомца
    addPetBtn.addEventListener('click', function (e) {
        e.preventDefault(); // Предотвращаем переход по ссылке
        showAddPetModal();
    });

    // Функция загрузки питомцев из localStorage
    function loadPetsFromStorage() {
        const savedPets = JSON.parse(localStorage.getItem('pets')) || [];

        if (savedPets.length > 0) {
            // Удаляем сообщение о пустом списке, если оно есть
            const emptyMessage = petsList.querySelector('.empty-pets-message');
            if (emptyMessage) petsList.removeChild(emptyMessage);

            // Добавляем каждого питомца из хранилища
            savedPets.forEach(pet => {
                addPetToDOM(pet);
            });
        }
    }

    // Функция добавления питомца в DOM
    function addPetToDOM(pet) {
        // Удаляем сообщение о пустом списке, если оно есть
        const emptyMessage = petsList.querySelector('.empty-pets-message');
        if (emptyMessage) petsList.removeChild(emptyMessage);

        // Создаем карточку питомца
        const petCard = document.createElement('div');
        petCard.className = 'pet-card';
        petCard.dataset.id = pet.id;
        petCard.onclick = function () { window.location.href = '/pet/' + pet.id; };

        petCard.innerHTML = `
            <div class="pet-icon-circle">
                <img src="${pet.imageUrl || './assets/img/free-icon-pets-12452249.png'}" alt="${pet.type}">
            </div>
            <div class="pet-info">
                <p><strong>Кличка:</strong> ${pet.name}</p>
                <p><strong>Вид:</strong> ${pet.type}</p>
                <p><strong>Порода:</strong> ${pet.breed}</p>
                <p><strong>Чип:</strong> ${pet.chip}</p>
                <p><strong>Возраст:</strong> ${pet.age}</p>
            </div>
            <button class="pet-profile-btn">Профиль</button>
            <button class="delete-pet-btn">Удалить</button>
        `;

        // Добавляем обработчик удаления
        const deleteBtn = petCard.querySelector('.delete-pet-btn');
        deleteBtn.addEventListener('click', function (e) {
            e.stopPropagation(); // Предотвращаем срабатывание клика по карточке
            deletePet(pet.id);
        });

        petsList.appendChild(petCard);
    }

    // Функция удаления питомца
    function deletePet(petId) {
        if (confirm('Вы уверены, что хотите удалить этого питомца?')) {
            // Удаляем из DOM
            const petCard = document.querySelector(`.pet-card[data-id="${petId}"]`);
            if (petCard) petCard.remove();

            // Удаляем из localStorage
            const savedPets = JSON.parse(localStorage.getItem('pets')) || [];
            const updatedPets = savedPets.filter(pet => pet.id !== petId);
            localStorage.setItem('pets', JSON.stringify(updatedPets));

            // Если питомцев не осталось, показываем сообщение
            if (updatedPets.length === 0) {
                showEmptyMessage();
            }
        }
    }

    // Функция показа модального окна добавления питомца
    function showAddPetModal() {
        // Создаем модальное окно
        const modal = document.createElement('div');
        modal.className = 'modal';
        modal.innerHTML = `
            <div class="modal-content">
                <span class="close-modal">&times;</span>
                <h2>Добавить нового питомца</h2>
                <form id="addPetForm">
                    <div class="form-group">
                        <label for="petName">Кличка:</label>
                        <input type="text" id="petName" required>
                    </div>
                    <div class="form-group">
                        <label for="petType">Вид:</label>
                        <select id="petType" required>
                            <option value="Кошка">Кошка</option>
                            <option value="Собака">Собака</option>
                        </select>
                    </div>
                    <div class="form-group">
                        <label for="petBreed">Порода:</label>
                        <input type="text" id="petBreed" required>
                    </div>
                    <div class="form-group">
                        <label for="petChip">Номер чипа:</label>
                        <input type="text" id="petChip">
                    </div>
                    <div class="form-group">
                        <label for="petAge">Возраст:</label>
                        <input type="text" id="petAge" required>
                    </div>
                    <div class="form-group">
                        <label for="petImage">Ссылка на фото (необязательно):</label>
                        <input type="text" id="petImage">
                    </div>
                    <button type="submit" class="submit-btn">Добавить питомца</button>
                </form>
            </div>
        `;

        document.body.appendChild(modal);

        // Обработчик закрытия модального окна
        const closeBtn = modal.querySelector('.close-modal');
        closeBtn.addEventListener('click', function () {
            document.body.removeChild(modal);
        });

        // Обработчик отправки формы
        const form = modal.querySelector('#addPetForm');
        form.addEventListener('submit', function (e) {
            e.preventDefault();
            handleAddPetForm();
        });
    }

    // Обработчик формы добавления питомца
    function handleAddPetForm() {
        // Получаем данные из формы
        const petData = {
            id: Date.now().toString(),
            name: document.getElementById('petName').value,
            type: document.getElementById('petType').value,
            breed: document.getElementById('petBreed').value,
            chip: document.getElementById('petChip').value,
            age: document.getElementById('petAge').value,
            imageUrl: document.getElementById('petImage').value || null
        };

        // Добавляем питомца
        addNewPet(petData);

        // Закрываем модальное окно
        const modal = document.querySelector('.modal');
        if (modal) document.body.removeChild(modal);
    }

    // Функция добавления нового питомца
    function addNewPet(petData) {
        // Добавляем в DOM
        addPetToDOM(petData);

        // Сохраняем в localStorage
        const savedPets = JSON.parse(localStorage.getItem('pets')) || [];
        savedPets.push(petData);
        localStorage.setItem('pets', JSON.stringify(savedPets));
    }

  
});
