// Handle answer option clicks
function initQuestionOptions() {
    const answersContainer = document.querySelector('.answers-container');
    if (!answersContainer) return;
    
    const questionId = parseInt(answersContainer.getAttribute('data-question-id'));
    const answerOptions = answersContainer.querySelectorAll('.answer-option.clickable');
    
    answerOptions.forEach(option => {
        option.addEventListener('click', function(e) {
            // Prevent multiple clicks
            if (this.classList.contains('submitting')) {
                return;
            }
            
            // Disable all options to prevent multiple submissions
            answerOptions.forEach(opt => {
                opt.classList.add('submitting');
                opt.style.pointerEvents = 'none';
                opt.style.opacity = '0.6';
            });
            
            const userAnswerId = parseInt(this.getAttribute('data-option-id'));
            
            fetch('/questions', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    question_id: questionId,
                    user_answer_id: userAnswerId
                })
            })
            .then(response => {
                if (response.redirected) {
                    window.location.href = response.url;
                } else {
                    return response.text().then(html => {
                        document.open();
                        document.write(html);
                        document.close();
                    });
                }
            })
            .catch(error => {
                console.error('Error:', error);
                alert('An error occurred while submitting your answer');
                // Re-enable options on error
                answerOptions.forEach(opt => {
                    opt.classList.remove('submitting');
                    opt.style.pointerEvents = 'auto';
                    opt.style.opacity = '1';
                });
            });
        });
    });
}

// Handle countdown timer for result page
function initCountdown() {
    const countdownElement = document.getElementById('countdown');
    if (!countdownElement) return;
    
    let countdown = 30;
    const interval = setInterval(() => {
        countdown--;
        countdownElement.textContent = countdown;
        if (countdown <= 0) {
            clearInterval(interval);
            window.location.href = '/question';
        }
    }, 1000);
}

// Initialize when DOM is ready
document.addEventListener('DOMContentLoaded', function() {
    initQuestionOptions();
    initCountdown();
});

