# Анимированные карточки (Animated Cards)

## Назначение

Этот пакет компонентов предоставляет реализацию анимированных карточек для создания интерактивных списков элементов с плавными анимациями появления и эффектами наведения.

## Компоненты

### 1. AnimatedCardList

Компонент для отображения списка анимированных карточек с поддержкой загрузки данных и индикаторов состояния.

#### Импорт

```javascript
import AnimatedCardList from './components/AnimatedCardList.vue'
```

#### Свойства (Props)

| Свойство | Тип | По умолчанию | Описание |
|----------|-----|--------------|----------|
| items | Array | [] | Массив элементов для отображения |
| loading | Boolean | false | Флаг состояния загрузки |
| noMoreItems | Boolean | false | Флаг отсутствия дополнительных элементов для загрузки |
| keyField | String | 'id' | Поле элемента, используемое в качестве ключа |
| cardClass | String | '' | Дополнительный CSS-класс для карточек |
| animationDelay | Number | 0.1 | Базовая задержка анимации в секундах |
| loadingText | String | 'Загрузка...' | Текст индикатора загрузки |
| noMoreItemsText | String | 'Все элементы загружены' | Текст сообщения об отсутствии элементов |

#### Слоты

- `header` - Слот для содержимого заголовка списка
- `card` - Слот для содержимого каждой карточки (получает параметры `item` и `index`)
- `footer` - Слот для содержимого подвала списка

#### Пример использования

```vue
<template>
  <AnimatedCardList 
    :items="users" 
    :loading="loading" 
    :no-more-items="noMoreUsers"
    key-field="id"
    card-class="user-card"
    :animation-delay="0.1"
    loading-text="Загрузка пользователей..."
    no-more-items-text="Все пользователи загружены"
  >
    <template #card="{ item: user }">
      <div class="user-header">
        <h3>{{ user.name }}</h3>
      </div>
      <div class="user-details">
        <p>{{ user.email }}</p>
      </div>
    </template>
  </AnimatedCardList>
</template>

<script setup>
import AnimatedCardList from './components/AnimatedCardList.vue'
import { ref } from 'vue'

const users = ref([
  { id: 1, name: 'Иван Иванов', email: 'ivan@example.com' },
  { id: 2, name: 'Мария Петрова', email: 'maria@example.com' }
])
const loading = ref(false)
const noMoreUsers = ref(true)
</script>
```

### 2. AnimatedCard

Компонент отдельной анимированной карточки.

#### Импорт

```javascript
import AnimatedCard from './components/AnimatedCard.vue'
```

#### Свойства (Props)

| Свойство | Тип | По умолчанию | Описание |
|----------|-----|--------------|----------|
| index | Number | 0 | Индекс элемента в списке (для вычисления задержки анимации) |
| animationDelay | Number | 0.1 | Базовая задержка анимации в секундах |
| customClass | String | '' | Дополнительный CSS-класс |
| itemId | [String, Number] | null | Уникальный идентификатор элемента |

#### Слоты

- `default` - Содержимое карточки

#### Пример использования

```vue
<template>
  <div class="cards-container">
    <AnimatedCard 
      v-for="(item, index) in items" 
      :key="item.id"
      :index="index"
      :animation-delay="0.1"
      custom-class="my-card"
      :item-id="item.id"
    >
      <h3>{{ item.title }}</h3>
      <p>{{ item.description }}</p>
    </AnimatedCard>
  </div>
</template>

<script setup>
import AnimatedCard from './components/AnimatedCard.vue'
import { ref } from 'vue'

const items = ref([
  { id: 1, title: 'Карточка 1', description: 'Описание карточки 1' },
  { id: 2, title: 'Карточка 2', description: 'Описание карточки 2' }
])
</script>
```

## Особенности поведения анимации

1. **Анимация появления**: Каждая карточка появляется снизу вверх с плавным увеличением прозрачности.
2. **Задержка анимации**: Задержка анимации вычисляется как `базовая_задержка + (индекс_элемента * 0.1)` секунд.
3. **Эффекты наведения**: При наведении карточки немного приподнимаются и усиливается тень.
4. **Поведение при добавлении элементов**: При добавлении новых элементов в уже отображенный список:
   - Новые элементы анимированно появляются с задержкой
   - Уже существующие элементы не перезапускают анимацию

## Стилизация

Компоненты используют CSS-переменные Telegram для адаптации к теме приложения:

- `--tg-theme-secondary-bg-color` - цвет фона карточек
- `--tg-theme-text-color` - цвет текста
- `--tg-theme-hint-color` - цвет вспомогательного текста
- `--tg-theme-button-color` - цвет акцентных элементов

Для кастомизации стилей можно переопределить CSS-классы:
- `.animated-card` - базовые стили карточки
- `.animated-card:hover` - стили при наведении