<template>
    <div class="headerBusy">
        <div class="headerBusy__title title">
            <div class="title__name">Токио</div>
            <div class="title__time time">
                <div class="time__col-1">
                    <img class="time__icon" src="@\assets\clock.svg" alt="" />
                    <span>11:14</span>
                </div>
                <div class="time__col-2">
                    <small>пт, 15 июля</small>
                </div>
            </div>
        </div>
        <div class="headerBusy__info info">
            <div class="info__equipment-list">
                <div class="info__equipment-item">
                    <img src="@\assets\Projector.svg" alt="icon" />
                    Проектор
                </div>
                <div class="info__equipment-item">
                    <img src="@\assets\TV.svg" alt="icon" />
                    Телевизор
                </div>
                <div class="info__equipment-item">
                    <img src="@\assets\Flipchart.svg" alt="icon" />
                    Флипчарт
                </div>
                <div class="info__equipment-item">
                    <img src="@\assets\board.svg" alt="icon" />
                    Доска для рисования
                </div>
                <div class="info__equipment-item">
                    <img src="@\assets\Terminal.svg" alt="icon" />
                    ВКС терминал
                </div>
            </div>
            <div class="info__lector lector">
                <div class="lector__foto">
                    <img src="@\assets\foto.svg" alt="lectors foto" />
                </div>
                <div class="lector__info">
                    <div class="lector__name">Иванова Анастасия</div>
                    <div class="lector__topic">Поведенческая стратегия в продажах</div>
                    <div class="lector__timeline">
                        <img src="@\assets\Progress.svg" alt="progress" />
                    </div>
                    <div class="lector__time">
                        <span>10:00</span>
                        <span>11:30</span>
                    </div>
                </div>
            </div>
        </div>
        <div class="headerBusy__footer">
            <button>Забронировать</button>
            <router-link to="/timeline" class="headerBusy__all-rooms all-rooms">
                <div class="all-rooms__text">
                    Другие переговорные
                    <div class="all-rooms__line"></div>
                </div>
                <img src="@\assets\arrow.svg" alt="arrow" />
            </router-link>
        </div>
        <div class="headerBusy__qr-code qr-code">
            <div class="qr-code__img-container">
                <img src="@\assets\qr.svg" alt="qr-code" />
            </div>
            <div class="qr-code__text">Забронируй с телефона</div>
        </div>
    </div>
</template>

<script>
export default {
    name: 'HeaderBusy',
    data() {
        return {};
    },
    methods: {},
};
</script>

<style scoped lang="scss">
.headerBusy {
    position: relative;
    display: flex;

    flex-direction: column;
    height: 100%;
    min-height: 336px;
    background: url(@/assets/HeaderBgBusy.png) no-repeat;
    background-size: cover;
    background-position: 50% 50%;
    mix-blend-mode: normal;
    border-radius: 12px;
    padding: 32px 24px 32px 32px;
    margin-bottom: 24px;
    &__title {
        display: flex;
        justify-content: space-between;
    }

    &__info {
        display: flex;
        flex-direction: row;
        flex: 1 0 auto;
        align-items: flex-start;
    }

    &__footer {
        display: flex;
        flex-direction: row;
        align-items: center;
        flex: 0 0 auto;
        button {
            font-weight: 600;
            font-size: 16px;
            line-height: 22px;
            width: 164px;
            height: 48px;
            background: #e11b11;
            border-radius: 12px;
            border: none;
            color: white;
        }
    }
    &__qr-code {
        position: absolute;
        padding: 16px;
        width: 92px;
        height: 121px;
        background: white;
        border-radius: 8px;
        right: 16px;
        bottom: 16px;
    }

    &__all-rooms {
        text-decoration: none;
        margin-left: 24px;
        display: flex;
        align-items: center;
        img {
            margin-left: 11px;
        }
    }
}
.title {
    font-weight: 700;
    font-size: 28px;
    line-height: 38px;

    &__time {
        display: flex;
        flex-direction: column;
        align-items: flex-end;
    }
}
.time {
    &__col-1 {
        display: flex;
        align-items: center;
    }

    &__icon {
        margin-right: 16px;
    }

    &__col-2 {
        font-weight: 600;
        font-size: 14px;
        line-height: 20px;
    }
}
.info {
    &__equipment-list {
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        max-width: 260px;
    }
    &__equipment-item {
        display: flex;
        align-items: center;
        font-weight: 600;
        font-size: 14px;
        line-height: 20px;
        margin-bottom: 9px;
        min-width: 120px;
        img {
            margin-right: 12px;
        }
    }

    &__lector {
        display: flex;
        flex-direction: row;
        align-items: center;
        margin-left: 37px;
    }
}
.lector {
    &__foto {
        margin-right: 14px;
        img {
            width: 92px;
            height: 92px;
            border-radius: 50%;
        }
    }

    &__name {
        font-weight: 700;
        font-size: 24px;
        line-height: 32px;
    }

    &__topic {
        font-weight: 600;
        font-size: 16px;
        line-height: 22px;
        margin-bottom: 15px;
    }

    &__timeline {
        margin-bottom: 8px;
    }

    &__time {
        font-weight: 700;
        font-size: 16px;
        line-height: 22px;
        display: flex;
        justify-content: space-between;
    }
}
.all-rooms {
    &__text {
        display: flex;
        flex-direction: column;
        font-weight: 600;
        font-size: 14px;
        line-height: 20px;
        color: white;
        outline: none;
        border: none;
    }
    &__line {
        border-bottom: 1px dashed #ffffff;
        align-self: stretch;
        opacity: 0.3;
    }
}
.qr-code {
    &__img-container {
        margin-bottom: 8px;
        img {
            width: 100%;
        }
    }
    &__text {
        color: #667387;
        font-weight: 400;
        font-size: 10px;
        line-height: 14px;
        text-align: center;
    }
}
</style>
