document.addEventListener('DOMContentLoaded', function() {
    const callBtn = document.querySelector('.call-btn'); // Кнопка вызова в основном интерфейсе
    const videoCallModal = document.getElementById('videoCallModal');
    const endCallBtn = document.getElementById('endCallBtn');
    const toggleMicBtn = document.getElementById('toggleMicBtn');
    const toggleCameraBtn = document.getElementById('toggleCameraBtn');
    const localVideo = document.getElementById('localVideo');
    const remoteVideo = document.getElementById('remoteVideo');
    
    let localStream;
    let isMicOn = true;
    let isCameraOn = true;

    // Имитация видеозвонка
    async function startCall() {
        try {
            localStream = await navigator.mediaDevices.getUserMedia({
                video: true,
                audio: true
            });
            
            localVideo.srcObject = localStream;
            
            // Имитация видео доктора
            setTimeout(() => {
                remoteVideo.srcObject = localStream.clone();
            }, 1500);
            
            videoCallModal.style.display = 'flex';
            
        } catch (error) {
            console.error('Ошибка доступа к медиаустройствам:', error);
            alert('Для видеозвонка необходимо разрешить доступ к камере и микрофону');
        }
    }

    function endCall() {
        if (localStream) {
            localStream.getTracks().forEach(track => track.stop());
        }
        if (remoteVideo.srcObject) {
            remoteVideo.srcObject.getTracks().forEach(track => track.stop());
        }
        
        localVideo.srcObject = null;
        remoteVideo.srcObject = null;
        videoCallModal.style.display = 'none';
        
        // Сброс состояний кнопок
        isMicOn = true;
        isCameraOn = true;
        toggleMicBtn.querySelector('img').src = 'assets/images/mic_on.png';
        toggleCameraBtn.querySelector('img').src = 'assets/images/camera_on.png';
    }

    function toggleMic() {
        if (localStream) {
            const audioTracks = localStream.getAudioTracks();
            audioTracks.forEach(track => {
                track.enabled = !track.enabled;
            });
            isMicOn = !isMicOn;
            toggleMicBtn.querySelector('img').src = isMicOn ? 
                'assets/images/mic_on.png' : 'assets/images/mic_off.png';
        }
    }

    function toggleCamera() {
        if (localStream) {
            const videoTracks = localStream.getVideoTracks();
            videoTracks.forEach(track => {
                track.enabled = !track.enabled;
            });
            isCameraOn = !isCameraOn;
            toggleCameraBtn.querySelector('img').src = isCameraOn ? 
                'assets/images/camera_on.png' : 'assets/images/camera_off.png';
        }
    }

    // Обработчики событий
    callBtn.addEventListener('click', startCall);
    endCallBtn.addEventListener('click', endCall);
    toggleMicBtn.addEventListener('click', toggleMic);
    toggleCameraBtn.addEventListener('click', toggleCamera);
});