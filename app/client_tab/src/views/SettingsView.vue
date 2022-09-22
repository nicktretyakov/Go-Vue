<!-- eslint-disable vuejs-accessibility/click-events-have-key-events -->
<!-- eslint-disable vuejs-accessibility/form-control-has-label -->
<template>
    <div class="SettingsView">
        <div class="SettingsView__title">Настройки планшета</div>
        <div class="SettingsView__inputs inputs">
            <div class="inputs__col1">
                <div class="inputs__col-line">
                    <span>Centrifuge URL</span>
                    <input required type="url" />
                </div>
                <div class="inputs__col-line">
                    <span>apiURL</span>
                    <input required type="url" />
                </div>
                <div class="inputs__col-line">
                    <span>API-key авторизации планшета в сервисе</span>
                    <input required type="text" value="" />
                </div>
            </div>
            <div class="inputs__col2">
                <div class="inputs__col-line">
                    <div class="blindCity"></div>
                    <span>Город</span>
                    <button class="drop-btn inputs__drop-city">
                        {{ btnIsActiveCity }}
                        <img src="@\assets\drop_arrow.svg" alt="" />
                    </button>
                    <div class="dropdown-content inputs__dropdown-city">
                        <button @click="activeCity" class="choose-btn" id="choose-btn">
                            Москва
                        </button>
                        <button @click="activeCity" class="choose-btn">Токио</button>
                        <button @click="activeCity" class="choose-btn">Пекин</button>
                        <button @click="activeCity" class="choose-btn">Минск</button>
                    </div>
                </div>
                <div class="inputs__col-line">
                    <div class="blindBuilding"></div>
                    <span>Здание</span>
                    <button class="drop-btn inputs__drop-building">
                        {{ btnIsActiveBuilding }}
                        <img src="@\assets\drop_arrow.svg" alt="" />
                    </button>
                    <div class="dropdown-content inputs__dropdown-building">
                        <button @click="activeBuilding" class="choose-btn">Гиляровского</button>
                        <button @click="activeBuilding" class="choose-btn">Лубянка</button>
                        <button @click="activeBuilding" class="choose-btn">Арбат</button>
                        <button @click="activeBuilding" class="choose-btn">Мира</button>
                    </div>
                </div>
                <div class="inputs__col-line">
                    <div class="blindRoom"></div>
                    <span>Переговорка</span>
                    <button class="drop-btn inputs__drop-room">
                        {{ btnIsActiveRoom }}
                        <img src="@\assets\drop_arrow.svg" alt="" />
                    </button>
                    <div class="dropdown-content inputs__dropdown-room">
                        <button @click="activeRoom" class="choose-btn">Китай</button>
                        <button @click="activeRoom" class="choose-btn">Беларусь</button>
                        <button @click="activeRoom" class="choose-btn">Индия</button>
                        <button @click="activeRoom" class="choose-btn">Литва</button>
                    </div>
                </div>
            </div>
        </div>
        <button disabled type="submit" class="SettingsView__save">Сохранить</button>
        <button @click="showPopup" class="SettingsView__default">Сбросить конфигурацию</button>
        <div class="SettingsView__version">TN Life: Версия 1.8.3 (build: 1347)</div>
        <div tag="div" name="fade" class="SettingsView__popup-wrapper popup-wrapper">
            <div class="popup-wrapper__popup">
                <div class="popup-wrapper__header">
                    Ты действительно хочешь сбросить конфигурацию?
                    <div @click="closePopup" class="popup-wrapper__cls-btn">
                        <img src="@\assets\cls-btn.svg" alt="cls-btn" />
                    </div>
                </div>
                <div class="popup-wrapper__btns">
                    <button @click="closePopup" class="popup-wrapper__cancel">Отмена</button>
                    <button class="popup-wrapper__confirm">Сбросить</button>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: 'SettingsView',
    data() {
        return {
            btnIsActiveCity: 'Выберите город',
            btnIsActiveBuilding: 'Выберите здание',
            btnIsActiveRoom: 'Выберите переговорку',
        };
    },
    methods: {
        activeCity(event) {
            this.btnIsActiveCity = event.target.textContent;
            const e = event.target;
            e.parentElement.style.display = 'none';
            e.parentElement.previousSibling.style.color = '#1E2228';
            document.querySelector('.blindBuilding').style.display = 'none';
            document.querySelector('.blindRoom').style.display = 'block';
            document.querySelector('.SettingsView__save').setAttribute('disabled', true);
            document.querySelector('.SettingsView__save').classList.remove('active');
            setTimeout(() => {
                e.parentElement.style.display = '';
            }, 500);
        },
        activeBuilding(event) {
            this.btnIsActiveBuilding = event.target.textContent;
            const e = event.target;
            e.parentElement.style.display = 'none';
            e.parentElement.previousSibling.style.color = '#1E2228';
            document.querySelector('.blindRoom').style.display = 'none';
            document.querySelector('.SettingsView__save').setAttribute('disabled', true);
            document.querySelector('.SettingsView__save').classList.remove('active');
            setTimeout(() => {
                e.parentElement.style.display = '';
            }, 500);
        },
        activeRoom(event) {
            this.btnIsActiveRoom = event.target.textContent;
            const e = event.target;
            e.parentElement.style.display = 'none';
            e.parentElement.previousSibling.style.color = '#1E2228';
            document.querySelector('.SettingsView__save').setAttribute('disabled', false);
            document.querySelector('.SettingsView__save').classList.add('active');
            setTimeout(() => {
                e.parentElement.style.display = '';
            }, 500);
        },
        showPopup() {
            document.querySelector('.popup-wrapper').style.visibility = 'visible';
        },
        closePopup() {
            document.querySelector('.popup-wrapper').style.visibility = 'hidden';
        },
    },
    mounted() {
        document.addEventListener('click', (item) => {
            if (item.target === document.querySelector('.popup-wrapper')) {
                document.querySelector('.popup-wrapper').style.visibility = 'hidden';
            }
        });
    },
};
</script>

<style lang="scss" scoped>
.SettingsView {
    font-weight: 600;
    font-size: 16px;
    line-height: 22px;
    color: #1e2228;
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 73px 40px 0;
    background: white;
    width: 100%;
    height: 100vh;

    &__title {
        font-weight: 700;
        font-size: 24px;
        line-height: 32px;
        margin-bottom: 54px;
    }

    &__inputs {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        margin-bottom: 66px;
        width: 100%;
    }

    &__save {
        font-weight: 700;
        font-size: 20px;
        line-height: 28px;
        color: #bcc1ce;
        width: 280px;
        min-height: 48px;
        background: #f9f9fa;
        border-radius: 16px;
        margin-bottom: 24px;
        cursor: default;
    }

    &__default {
        width: 280px;
        min-height: 48px;
        background: #e6e8ed;
        border-radius: 12px;
        margin-bottom: 20px;
    }
    &__version {
        font-weight: 400;
        font-size: 16px;
        line-height: 22px;
        color: #bcc1ce;
        margin: auto auto 25px auto;
    }
    &__popup-wrapper {
        visibility: hidden;
        position: fixed;
        top: 0;
        right: 0;
        bottom: 0;
        left: 0;
        background: rgb(30 34 40 / 40%);
        z-index: 3;
        display: flex;
        justify-content: center;
        align-items: center;
    }
}
.active {
    background: #e11b11;
    color: #ffffff;
    cursor: pointer;
}
.inputs {
    &__col1 {
        display: flex;
        flex-direction: column;
        width: 100%;
        margin-right: 30px;
    }
    &__col2 {
        display: flex;
        flex-direction: column;
        width: 100%;
        margin-left: 30px;
    }

    &__col-line {
        position: relative;
        display: flex;
        flex-direction: column;
    }
    &__drop-city:hover + &__dropdown-city,
    &__dropdown-city:hover {
        display: flex;
    }
    &__drop-building:hover + &__dropdown-building,
    &__dropdown-building:hover {
        display: flex;
    }
    &__drop-room:hover + &__dropdown-room,
    &__dropdown-room:hover {
        display: flex;
    }
}
input,
.drop-btn {
    width: 100%;
    height: 48px;
    border: 1px solid #d7dae1;
    border-radius: 10px;
    background: #ffffff;
    margin-bottom: 24px;
    padding: 12px 16px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    font-weight: 400;
    font-size: 16px;
    line-height: 22px;
    color: #1e2228;
}
span {
    margin-bottom: 10px;
}
.drop-btn {
    color: #667387;
}
.blindBuilding,
.blindRoom,
.blind {
    width: 100%;
    height: 49px;
    position: absolute;
    top: 31px;
    background: #d7dae1;
    border-radius: 10px;
    opacity: 0.4;
    z-index: 1;
}
.dropdown-content {
    width: 100%;
    flex-direction: column;
    position: absolute;
    z-index: 2;
    background: white;
    border-radius: 10px;
    top: 80px;
    box-shadow: 0px 0px 2px rgba(30, 34, 40, 0.19), -4px 6px 7px rgba(30, 34, 40, 0.17);
    border-radius: 12px;
    display: none;
}
.choose-btn {
    height: 48px;
}
.popup-wrapper {
    &__popup {
        width: 466px;
        height: 200px;
        background: #ffffff;
        border-radius: 16px;
        padding: 28px 24px 32px;
    }

    &__header {
        font-weight: 700;
        font-size: 24px;
        line-height: 32px;
        color: #1e2228;
        margin-bottom: 28px;
        display: flex;
        flex-direction: row;
        justify-content: space-between;
    }

    &__btns {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        font-weight: 600;
        font-size: 16px;
        line-height: 22px;
    }

    &__cancel {
        width: 189px;
        height: 48px;
        background: #ffffff;
        border: 1px solid #e11b11;
        border-radius: 12px;
        color: #e11b11;
    }

    &__confirm {
        width: 209px;
        height: 48px;
        background: #e11b11;
        border-radius: 12px;
        color: #ffffff;
    }
    &__cls-btn {
        margin-left: 30px;
        margin-right: 13px;
        cursor: pointer;
        img {
            width: 14px;
        }
    }
}
</style>
