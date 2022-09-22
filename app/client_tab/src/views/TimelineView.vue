<template>
    <div class="timeline">
        <div class="timeline__title">Все переговорные</div>
        <div class="timeline__content">
            <div class="timeline__location location">
                <div class="location__item">
                    <div class="location__icon">
                        <img src="@\assets\icon-city.svg" alt="icon-city" />
                    </div>
                    <div class="location__info info">
                        <div class="info__city">Монако</div>
                        <div class="info__space">2 этаж <span></span> 3 места</div>
                    </div>
                    <div class="location__arrow">
                        <img src="@\assets\location_arrow.svg" alt="" />
                    </div>
                </div>
                <div class="location__item">
                    <div class="location__icon">
                        <img src="@\assets\icon-city.svg" alt="icon-city" />
                    </div>
                    <div class="location__info info">
                        <div class="info__city">Беларусь</div>
                        <div class="info__space">2 этаж <span></span> 8 мест</div>
                    </div>
                    <div class="location__arrow">
                        <img src="@\assets\location_arrow.svg" alt="" />
                    </div>
                </div>
                <div class="location__item">
                    <div class="location__icon">
                        <img src="@\assets\icon-city.svg" alt="icon-city" />
                    </div>
                    <div class="location__info info">
                        <div class="info__city">Беларусь</div>
                        <div class="info__space">2 этаж <span></span> 8 мест</div>
                    </div>
                    <div class="location__arrow">
                        <img src="@\assets\location_arrow.svg" alt="" />
                    </div>
                </div>
                <div class="location__item">
                    <div class="location__icon">
                        <img src="@\assets\icon-city.svg" alt="icon-city" />
                    </div>
                    <div class="location__info info">
                        <div class="info__city">Беларусь</div>
                        <div class="info__space">2 этаж <span></span> 8 мест</div>
                    </div>
                    <div class="location__arrow">
                        <img src="@\assets\location_arrow.svg" alt="" />
                    </div>
                </div>
                <div class="location__item">
                    <div class="location__icon">
                        <img src="@\assets\icon-city.svg" alt="icon-city" />
                    </div>
                    <div class="location__info info">
                        <div class="info__city">Беларусь</div>
                        <div class="info__space">2 этаж <span></span> 8 мест</div>
                    </div>
                    <div class="location__arrow">
                        <img src="@\assets\location_arrow.svg" alt="" />
                    </div>
                </div>
            </div>
            <div class="timeline__wrapper">
                <ul class="timeline__time">
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
                <div class="timeline__col">
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
                <div class="timeline__rt-line" style="left: 25%"></div>
            </div>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.timeline {
    padding: 32px 24px;
    display: flex;
    flex-direction: column;
    background: white;
    height: 100%;
    &__title {
        color: #1e2228;
        font-weight: 700;
        font-size: 24px;
        line-height: 32px;
        margin-bottom: 14px;
    }

    &__content {
        display: flex;
        flex-direction: row;
        height: 100%;
    }

    &__location {
        margin-top: 40px;
        width: 100%;
        max-width: 256px;
    }

    &__wrapper {
        width: 100%;
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
        height: 100%;
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
.location {
    &__item {
        width: 100%;
        height: 70px;
        border: 1px solid #d7dae1;
        border-radius: 16px;
        padding: 14px 28px 14px 14px;
        display: flex;
        flex-direction: row;
        align-items: center;
        margin-bottom: 14px;
    }

    &__icon {
        width: 42px;
        height: 42px;
        img {
            width: 100%;
            height: auto;
        }
    }

    &__info {
        margin-left: 12px;
    }

    &__arrow {
        margin-left: auto;
    }
}
.info {
    &__city {
        font-weight: 400;
        font-size: 16px;
        line-height: 22px;
        color: black;
    }

    &__space {
        font-weight: 400;
        font-size: 12px;
        line-height: 16px;
        color: #667387;
        span {
            display: inline-block;
            vertical-align: middle;
            width: 2px;
            height: 2px;
            border-radius: 50%;
            border: 2px solid #667387;
            margin: 0 5px;
            background: #667387;
        }
    }
}
</style>
