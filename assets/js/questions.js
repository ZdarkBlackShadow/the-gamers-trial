// Handle question form submission
function initQuestionForm() {
    const questionForm = document.getElementById('question-form');
    if (!questionForm) return;
    
    questionForm.addEventListener('submit', function(e) {
        e.preventDefault();
        const questionId = parseInt(document.getElementById('question_id').value);
        const selectedAnswer = document.querySelector('input[name="user_answer_id"]:checked');
        
        if (!selectedAnswer) {
            alert('Please select an answer');
            return;
        }
        
        const userAnswerId = parseInt(selectedAnswer.value);
        
        fetch('/question', {
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
    initQuestionForm();
    initCountdown();
});

