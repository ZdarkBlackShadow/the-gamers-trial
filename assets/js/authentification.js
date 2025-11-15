function switchTab(tab) {
    // Masquer tous les formulaires
    document.querySelectorAll('.auth-form').forEach(form => {
        form.classList.remove('active');
    });
    
    // Retirer la classe active de tous les boutons
    document.querySelectorAll('.tab-btn').forEach(btn => {
        btn.classList.remove('active');
    });
    
    // Afficher le formulaire sélectionné
    if (tab === 'register') {
        document.getElementById('register-form').classList.add('active');
        document.querySelectorAll('.tab-btn')[0].classList.add('active');
    } else {
        document.getElementById('login-form').classList.add('active');
        document.querySelectorAll('.tab-btn')[1].classList.add('active');
    }
}

// Initialisation quand le DOM est chargé
document.addEventListener('DOMContentLoaded', function() {
    // Ajouter les event listeners aux boutons d'onglets
    const registerTab = document.querySelectorAll('.tab-btn')[0];
    const loginTab = document.querySelectorAll('.tab-btn')[1];
    
    if (registerTab) {
        registerTab.addEventListener('click', function() {
            switchTab('register');
        });
    }
    
    if (loginTab) {
        loginTab.addEventListener('click', function() {
            switchTab('login');
        });
    }

    // Gestion de la soumission du formulaire d'inscription
    const registerForm = document.getElementById('registerForm');
    if (registerForm) {
        registerForm.addEventListener('submit', function(e) {
            e.preventDefault();
            const formData = {
                email: document.getElementById('register-email').value,
                pseudo: document.getElementById('register-pseudo').value,
                password: document.getElementById('register-password').value,
                password_confirmation: document.getElementById('register-password-confirm').value
            };
            
            fetch('/authentification/register', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(formData),
                redirect: 'follow'
            })
            .then(response => {
                // Le backend redirige toujours (succès vers /home, erreur vers /authentification)
                if (response.redirected) {
                    window.location.href = response.url;
                } else if (response.ok) {
                    // Si pas de redirection mais OK, on redirige quand même
                    window.location.href = '/home';
                } else {
                    // En cas d'erreur, on recharge la page pour afficher l'erreur depuis la session
                    window.location.href = '/authentification';
                }
            })
            .catch(error => {
                console.error('Error:', error);
                // En cas d'erreur réseau, on recharge quand même pour voir l'erreur
                window.location.href = '/authentification';
            });
        });
    }

    // Gestion de la soumission du formulaire de connexion
    const loginForm = document.getElementById('loginForm');
    if (loginForm) {
        loginForm.addEventListener('submit', function(e) {
            e.preventDefault();
            const formData = {
                email_or_pseudo: document.getElementById('login-email-pseudo').value,
                password: document.getElementById('login-password').value
            };
            
            fetch('/authentification/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(formData),
                redirect: 'follow'
            })
            .then(response => {
                // Le backend redirige toujours (succès vers /home, erreur vers /authentification)
                if (response.redirected) {
                    window.location.href = response.url;
                } else if (response.ok) {
                    // Si pas de redirection mais OK, on redirige quand même
                    window.location.href = '/home';
                } else {
                    // En cas d'erreur, on recharge la page pour afficher l'erreur depuis la session
                    window.location.href = '/authentification';
                }
            })
            .catch(error => {
                console.error('Error:', error);
                // En cas d'erreur réseau, on recharge quand même pour voir l'erreur
                window.location.href = '/authentification';
            });
        });
    }
});

