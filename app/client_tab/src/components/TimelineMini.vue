<!-- eslint-disable vuejs-accessibility/click-events-have-key-events -->
<template>
    <div class="timelineMini">
        <div class="timelineMini__hidden">
            <div class="timelineMini__wrapper">
                <ul class="timelineMini__time">
                    <li>9:00</li>
                    <li>10:00</li>
                    <li>11:00</li>
                    <li>12:00</li>
                    <li>13:00</li>
                    <li>14:00</li>
                    <li>15:00</li>
                    <li>16:00</li>
                    <li>17:00</li>
                    <li>18:00</li>
                </ul>
                <div class="timelineMini__col">
                    <div></div>
                    <div></div>
                    <div></div>
                    <div></div>
                    <div></div>
                    <div></div>
                    <div></div>
                    <div></div>
                    <div></div>
                    <div></div>
                    <div></div>
                    <div></div>
                    <div></div>
                    <div></div>
                    <div></div>
                    <div></div>
                    <div></div>
                    <div></div>
                    <div></div>
                    <div></div>
                </div>
                <div class="timelineMini__items">
                    <div class="timelineMini__item" style="width: 52px; left: 0px">
                        <div class="timelineMini__foto">
                            <img src="@\assets\boroda.svg" alt="union" />
                        </div>
                        <div @click="showPopup" class="timelineMini__info">
                            <img src="@\assets\Union.svg" alt="union" />
                        </div>
                        <div class="timelineMini__popup-info popup-info">
                            <div class="popup-info__lectors-name">Иванова Анастасия</div>
                            <div class="popup-info__topic">Поведенческая стратегия в продажах</div>
                            <div class="popup-info__time">12:00 — 12:30</div>
                            <div @click="closePopup" class="popup-info__cls-btn">
                                <img src="@\assets\cls-btn.svg" alt="cls-btn" />
                            </div>
                        </div>
                    </div>
                    <div class="timelineMini__item" style="width: 157px; left: 107px">
                        <div class="timelineMini__foto">
                            <img src="@\assets\foto.svg" alt="union" />
                        </div>
                        <div @click="showPopup" class="timelineMini__info">
                            <img src="@\assets\Union.svg" alt="union" />
                        </div>
                        <div class="timelineMini__popup-info popup-info">
                            <div class="popup-info__lectors-name">Иванова Анастасия</div>
                            <div class="popup-info__topic">Поведенческая стратегия в продажах</div>
                            <div class="popup-info__time">12:00 — 12:30</div>
                            <div @click="closePopup" class="popup-info__cls-btn">
                                <img src="@\assets\cls-btn.svg" alt="cls-btn" />
                            </div>
                        </div>
                    </div>
                    <div class="timelineMini__item" style="width: 52px; left: 527px">
                        <div class="timelineMini__foto">ВА</div>
                        <div @click="showPopup" class="timelineMini__info">
                            <img src="@\assets\Union.svg" alt="union" />
                        </div>
                        <div class="timelineMini__popup-info popup-info">
                            <div class="popup-info__lectors-name">Иванова Анастасия</div>
                            <div class="popup-info__topic">Поведенческая стратегия в продажах</div>
                            <div class="popup-info__time">12:00 — 12:30</div>
                            <div @click="closePopup" class="popup-info__cls-btn">
                                <img src="@\assets\cls-btn.svg" alt="cls-btn" />
                            </div>
                        </div>
                    </div>
                </div>
                <div class="timelineMini__rt-line" style="left: 25%"></div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: 'TimelineMini',
    data() {
        return {};
    },
    methods: {
        closePopup(event) {
            const e = event.target;
            e.parentElement.parentElement.style.visibility = 'hidden';
        },
        showPopup(event) {
            const e = event.target.parentElement.nextSibling;
            e.style.visibility = 'visible';
            setTimeout(() => {
                e.style.visibility = 'hidden';
            }, 7000);
        },
        // createChart(e) {
        //     const days = document.querySelectorAll('.chart-values li');
        //     const tasks = document.querySelectorAll('.chart-bars li');
        //     const daysArray = [...days];
        //     tasks.forEach((el) => {
        //         const duration = el.dataset.duration.split('-');
        //         const startDay = duration[0];
        //         const endDay = duration[1];
        //         let left = 0;
        //         let width = 0;
        //         if (startDay.endsWith('½')) {
        //             const filteredArray = daysArray.filter(
        //                 (day) => day.textContent === startDay.slice(0, -1)
        //             );
        //             left = filteredArray[0].offsetLeft + filteredArray[0].offsetWidth / 2;
        //         } else {
        //             const filteredArray = daysArray.filter((day) => day.textContent === startDay)
        //             left = filteredArray[0].offsetLeft;
        //         }
        //         if (endDay.endsWith('½')) {
        //             const filteredArray = daysArray.filter(
        //                 (day) => day.textContent === endDay.slice(0, -1)
        //             );
        //             width = filteredArray[0].offsetLeft + filteredArray[0].offsetWidth / 2 - left
        //         } else {
        //             const filteredArray = daysArray.filter((day) => day.textContent === endDay);
        //             width = filteredArray[0].offsetLeft + filteredArray[0].offsetWidth - left;
        //         }
        //         // apply css
        //         el.style.left = `${left}px`;
        //         el.style.width = `${width}px`;
        //         if (e.type === 'load') {
        //             el.style.backgroundColor = el.dataset.color;
        //             el.style.opacity = 1;
        //         }
        //     });
        // },
    },
    // created() {
    //     window.addEventListener('load', this.createChart);
    //     window.addEventListener('resize', this.createChart);
    // },
};
</script>

<style scoped lang="scss">
.timelineMini {
    position: relative;
    height: 100%;
    min-height: 268px;
    background: white;
    border-radius: 16px;
    display: flex;
    flex: 1;
    padding: 42px 30px;
    &__hidden {
        display: flex;
        flex: 1;
        overflow: hidden;
    }

    &__wrapper {
        position: relative;
        display: flex;
        flex-direction: column;
        overflow: hidden;
        overflow-x: scroll;
    }
    &__time {
        font-weight: 700;
        font-size: 16px;
        line-height: 22px;
        color: #1e2228;
        flex: 1;
        display: flex;
        flex-direction: row;
        margin: 0;
        padding: 0;
        li {
            width: 104px;
            flex-shrink: 0;
            margin: 0;
        }
    }
    &__col {
        height: 280px;
        display: flex;
        flex-direction: row;
        margin-left: 10px;
        div {
            width: 52px;
            flex-shrink: 0;
        }
        div:nth-child(even) {
            border-right: 1px solid #d7dae1;
        }
        div:last-child {
            border-right: none;
        }
        div:first-child {
            border-left: 1px solid #d7dae1;
        }
    }
    &__items {
        height: 249px;
        width: 1030px;
        position: absolute;
        top: 30px;
        display: flex;
        flex-direction: row;
    }

    &__item {
        position: inherit;
        width: 50px;
        height: 100%;
        background: #e6e8ed;
        border-radius: 12px;
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        align-items: center;
        padding: 15px 0;
    }
    &__info {
        cursor: pointer;
    }
    &__popup-info {
        visibility: hidden;
        position: absolute;
        width: 235px;
        height: 115px;
        top: 0;
        left: 110%;
        padding: 14px 20px;
        background: #ffffff;
        box-shadow: 0px 18px 44px rgba(0, 0, 0, 0.09);
        border-radius: 14px;
        z-index: 2;
        display: flex;
        flex-direction: column;
    }

    &__foto {
        background: #e11b11;
        color: white;
        height: 32px;
        width: 32px;
        line-height: 32px;
        border-radius: 50%;
        overflow: hidden;
        text-align: center;
        img {
            height: 100%;
        }
    }
    &__rt-line {
        display: block;
        position: absolute;
        top: 18px;
        bottom: 0;
        width: 2px;
        height: 100%;
        background-color: #e11b11;
        box-shadow: 0px 3px 23px rgb(225 27 17 / 14%);
        z-index: 1;

        &::before {
            content: '';
            display: block;
            width: 12px;
            height: 12px;
            background-color: #e11b11;
            border: 2px solid #fff;
            margin-left: -4.5px;
            border-radius: 50%;
            box-shadow: 0px 3px 12px rgb(225 27 17 / 16%);
        }
    }
}
.popup-info {
    &__lectors-name {
        font-weight: 700;
        font-size: 16px;
        line-height: 22px;
        color: #1e2228;
        margin-bottom: 4px;
    }
    &__topic {
        font-weight: 400;
        font-size: 14px;
        line-height: 16px;
        color: #1e2228;
        margin-bottom: 12px;
    }
    &__time {
        font-weight: 400;
        font-size: 14px;
        line-height: 16px;
        color: #667387;
    }
    &__cls-btn {
        position: absolute;
        left: 90%;
        cursor: pointer;
    }
}
</style>
