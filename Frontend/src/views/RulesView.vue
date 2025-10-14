<template>
    <div class="rules-page">

        <main class="rules-content">
            <section class="game-description">
                <div class="intro-card">
                    <h2>Добро пожаловать в Эпическую Битву Кланов!</h2>
                    <p class="epic-description">
                        Готовы стать частью самой захватывающей социальной игры для больших компаний?
                        Prophecy - это не просто игра, это настоящее приключение, где каждый выбор имеет значение,
                        а каждая встреча может изменить ход битвы!
                    </p>
                    <div class="game-mechanics">
                        <h3>Как это работает?</h3>
                        <ul>
                            <li>🎮 Игра предназначена для 20+ участников на масштабных мероприятиях</li>
                            <li>🤝 Взаимодействие начинается только после сканирования QR-кодов новых знакомых</li>
                            <li>🎭 Роли распределяются случайным образом и могут меняться в процессе игры</li>
                            <li>🏰 Создавайте кланы, объединяйтесь в альянсы или становитесь одиночками-убийцами</li>
                            <li>⚡ Каждое действие влияет на баланс сил в игре</li>
                        </ul>
                    </div>
                </div>
            </section>

            <section class="roles-section">
                <div class="roles-header-card">
                    <h2>ЛЕГЕНДАРНЫЕ РОЛИ</h2>
                    <p class="roles-subtitle">Судьба выберет тебя - будь готов к любому испытанию!</p>
                </div>

                <div class="roles-grid">
                    <div v-for="(role, index) in roles" :key="role.name" class="role-card animated-card"
                        :style="{ animationDelay: `${index * 0.1}s` }">
                        <div class="role-header">
                            <span class="role-emoji">{{ role.emoji }}</span>
                            <h3 class="role-name">{{ role.name }}</h3>
                        </div>
                        <p class="role-description">{{ role.description }}</p>
                    </div>
                </div>
            </section>

            <section class="cta-section">
                <div class="cta-card">
                    <h2>ГОТОВЫ К ПРИКЛЮЧЕНИЯМ?</h2>
                    <p>Присоединяйтесь к битве и докажите, что вы достойны легенд!</p>
                    <button class="cta-button" @click="goToHome">Начать Игру</button>
                </div>
            </section>
        </main>
    </div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import HeaderNav from '@/components/HeaderNav.vue'

const router = useRouter()

const goToHome = () => {
    router.push('/')
}

const roles = [
    {
        name: "Мастер Кодекса",
        emoji: "👑",
        description: "Ты - Мастер Кодекса. Твоя миссия - собрать мощный клан, привлекая как можно больше участников. Ты получаешь поинты за каждого привлеченного участника и имеешь власть удалять людей из клана без ожидания. Ты - лидер, ты - власть."
    },
    {
        name: "Страж Правосудия",
        emoji: "👮",
        description: "Ты - Страж Правосудия. Твоя задача - охотиться на предателей в клане. Ты получаешь очки за каждого найденного предателя и можешь покинуть клан в любое время. Но будь осторожен, если ты ошибешься и обвинишь невиновного, создатель клана получит оповещение. Ты - закон, ты - справедливость."
    },
    {
        name: "Целитель Душ",
        emoji: "🩺",
        description: "Ты - Целитель Душ. Ты не можешь покинуть клан самостоятельно, но ты можешь исцелять его участников. Ты получаешь поинты за каждого вылеченного участника. Чтобы вылечить участника, тебе нужно добавить его в друзья. Ты - исцеление, ты - жизнь."
    },
    {
        name: "Теневой Манипулятор",
        emoji: "🎭",
        description: "Ты - Теневой Манипулятор. Твоя задача - уничтожать участников клана. Ты не можешь покинуть клан самостоятельно. У тебя есть время отката между убийствами. При каждом убийстве создатель клана получает оповещение, но твое имя остается в тайне. Ты - тень, ты - смерть."
    },
    {
        name: "Герой",
        emoji: "⚔️",
        description: "Ты - Герой. Ты - единственный в своем клане. Твоя задача - собрать войско, добавив как можно больше участников клана в друзья и пригласив их в войско. Ты можешь совершать набеги на территории других кланов. При победах все участники битвы получают очки, при поражении - теряют. Ты - мощь, ты - победа."
    },
    {
        name: "Неуловимый Убийца",
        emoji: "🗡",
        description: "Ты - Неуловимый Убийца. Тебе выдаётся цель для убийства. Ник цели известен, и цель получает оповещение. Попытка убийства происходит через 10 минут после добавления цели в друзья. Если цель не в друзьях, то происходит смена цели. Если убийство произошло удачно, ты получаешь поинты, а цель не может взаимодействовать с меню в течение 10 минут. Ты - страх, ты - неизбежность."
    },
    {
        name: "Непробиваемый Защитник",
        emoji: "🛡",
        description: "Ты - Непробиваемый Защитник. При появлении новой цели для убивца ты получаешь оповещение, включая ник. Ты должен добавиться в друзья и можешь проверить любого друга цели раз в минуту. Если проверка удачна, ты получаешь валюту. Ты - стена, ты - защита."
    },
    {
        name: "Сыщик",
        emoji: "🔍",
        description: "Ты - Сыщик. Твоя задача - вычислить роль другого участника. Для этого тебе нужно добавить его в друзья и выбрать роль. За каждый неправильный ответ время отката увеличивается, но за правильный ответ ты получаешь очки, а жертва теряет часть своих очков. Ты - знание, ты - мудрость."
    }
]
</script>

<style scoped>
.rules-page {
    min-height: 100vh;
    background-color: var(--tg-theme-bg-color, #f5f5f5);
    color: var(--tg-theme-text-color, #333333);
    padding-bottom: 40px;
    font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
    transition: all 0.3s ease;
}

.rules-content {
    max-width: 1200px;
    margin: 0 auto;
    padding: 20px;
}

.intro-card {
    background: var(--tg-theme-secondary-bg-color, white);
    border-radius: 20px;
    padding: 30px;
    box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
    border: 1px solid var(--tg-theme-secondary-bg-color, #e0e0e0);
    margin-bottom: 30px;
    transition: all 0.3s ease;
}

:global(.tg-theme-dark) .intro-card {
    box-shadow: 0 2px 12px rgba(0, 0, 0, 0.3);
}

.intro-card h2 {
    font-size: 2.2rem;
    margin-bottom: 20px;
    color: var(--tg-theme-text-color, #333333);
    text-align: center;
    transition: color 0.3s ease;
}

.epic-description {
    font-size: 1.2rem;
    line-height: 1.7;
    text-align: center;
    margin-bottom: 25px;
    color: var(--tg-theme-text-color, #333333);
    transition: color 0.3s ease;
}

.game-mechanics h3 {
    font-size: 1.5rem;
    margin-bottom: 15px;
    color: var(--tg-theme-text-color, #333333);
    text-align: center;
    transition: color 0.3s ease;
}

.game-mechanics ul {
    list-style-type: none;
    padding: 0;
    margin: 0;
}

.game-mechanics li {
    padding: 12px 20px;
    margin-bottom: 10px;
    background: var(--tg-theme-secondary-bg-color, white);
    border-radius: 10px;
    border-left: 4px solid var(--tg-theme-button-color, #667eea);
    font-size: 1.1rem;
    display: flex;
    align-items: center;
    transition: all 0.3s ease;
}

:global(.tg-theme-dark) .game-mechanics li {
    background: rgba(255, 255, 255, 0.1);
}

.game-mechanics li:hover {
    background: var(--tg-theme-button-color, #667eea);
    color: var(--tg-theme-button-text-color, white);
}

.game-mechanics li:before {
    content: "⚡";
    margin-right: 10px;
    font-size: 1.3rem;
}

.roles-header-card {
    background: var(--tg-theme-secondary-bg-color, white);
    border-radius: 20px;
    padding: 25px;
    box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
    border: 1px solid var(--tg-theme-secondary-bg-color, #e0e0e0);
    text-align: center;
    margin-bottom: 30px;
    transition: all 0.3s ease;
}

:global(.tg-theme-dark) .roles-header-card {
    box-shadow: 0 2px 12px rgba(0, 0, 0, 0.3);
}

.roles-header-card h2 {
    font-size: 2.5rem;
    margin-bottom: 10px;
    color: var(--tg-theme-text-color, #333333);
    transition: color 0.3s ease;
}

.roles-subtitle {
    font-size: 1.3rem;
    color: var(--tg-theme-hint-color, #666666);
    font-style: italic;
    transition: color 0.3s ease;
}

.roles-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
    gap: 25px;
    margin-bottom: 30px;
}

.role-card {
    background: var(--tg-theme-secondary-bg-color, white);
    border-radius: 20px;
    padding: 25px;
    box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
    border: 1px solid var(--tg-theme-secondary-bg-color, #e0e0e0);
    transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275);
    position: relative;
    overflow: hidden;
    opacity: 0;
    transform: translateY(20px);
    animation: fadeInUp 0.6s forwards;
}

:global(.tg-theme-dark) .role-card {
    box-shadow: 0 2px 12px rgba(0, 0, 0, 0.3);
}

.role-card:hover {
    transform: translateY(-5px);
    box-shadow: 0 8px 25px rgba(0, 0, 0, 0.15);
}

:global(.tg-theme-dark) .role-card:hover {
    box-shadow: 0 8px 25px rgba(0, 0, 0, 0.4);
}

.role-header {
    display: flex;
    align-items: center;
    margin-bottom: 20px;
    padding-bottom: 15px;
    border-bottom: 1px solid var(--tg-theme-hint-color, #e0e0e0);
}

.role-emoji {
    font-size: 2.5rem;
    margin-right: 15px;
}

.role-name {
    font-size: 1.6rem;
    margin: 0;
    color: var(--tg-theme-text-color, #333333);
    transition: color 0.3s ease;
}

.role-description {
    font-size: 1.05rem;
    line-height: 1.6;
    margin: 0;
    color: var(--tg-theme-text-color, #333333);
    transition: color 0.3s ease;
}

.cta-section {
    text-align: center;
}

.cta-card {
    background: var(--tg-theme-secondary-bg-color, white);
    border-radius: 20px;
    padding: 40px;
    box-shadow: 0 2px 12px rgba(0, 0, 0, 0.08);
    border: 1px solid var(--tg-theme-secondary-bg-color, #e0e0e0);
    transition: all 0.3s ease;
}

:global(.tg-theme-dark) .cta-card {
    box-shadow: 0 2px 12px rgba(0, 0, 0, 0.3);
}

.cta-card h2 {
    font-size: 2.2rem;
    margin-bottom: 15px;
    color: var(--tg-theme-text-color, #333333);
    transition: color 0.3s ease;
}

.cta-card p {
    font-size: 1.3rem;
    margin-bottom: 30px;
    color: var(--tg-theme-text-color, #333333);
    transition: color 0.3s ease;
}

.cta-button {
    background: var(--tg-theme-button-color, #667eea);
    color: var(--tg-theme-button-text-color, white);
    border: none;
    padding: 15px 40px;
    font-size: 1.2rem;
    border-radius: 50px;
    cursor: pointer;
    transition: all 0.3s ease;
    font-weight: bold;
    letter-spacing: 1px;
    text-transform: uppercase;
}

.cta-button:hover {
    opacity: 0.9;
    transform: translateY(-3px);
    box-shadow: 0 5px 15px rgba(0, 0, 0, 0.2);
}

@keyframes fadeInUp {
    to {
        opacity: 1;
        transform: translateY(0);
    }
}

@media (max-width: 768px) {
    .rules-content {
        padding: 15px;
    }

    .intro-card,
    .roles-header-card,
    .role-card,
    .cta-card {
        padding: 20px;
    }

    .intro-card h2 {
        font-size: 1.8rem;
    }

    .roles-header-card h2 {
        font-size: 2rem;
    }

    .roles-grid {
        grid-template-columns: 1fr;
    }

    .role-header {
        flex-direction: column;
        text-align: center;
    }

    .role-emoji {
        font-size: 3rem;
        margin-right: 0;
        margin-bottom: 10px;
    }

    .role-name {
        font-size: 1.4rem;
    }

    .game-mechanics li,
    .role-description {
        font-size: 1rem;
    }

    .cta-card h2 {
        font-size: 1.8rem;
    }

    .cta-card p {
        font-size: 1.1rem;
    }

    .cta-button {
        padding: 12px 30px;
        font-size: 1rem;
    }
}

@media (max-width: 480px) {
    .rules-content {
        padding: 10px;
    }

    .intro-card,
    .roles-header-card,
    .role-card,
    .cta-card {
        padding: 15px;
    }

    .intro-card h2,
    .roles-header-card h2 {
        font-size: 1.5rem;
    }

    .role-name {
        font-size: 1.2rem;
    }

    .role-emoji {
        font-size: 2.5rem;
    }

    .cta-card h2 {
        font-size: 1.5rem;
    }

    .cta-card p {
        font-size: 1rem;
    }
}
</style>