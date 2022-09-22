<template>
    <div class="home container">
        <HeaderFree />
        <HeaderBusy />
        <TimelineMini />
    </div>
</template>

<script>
import HeaderFree from '@/components/HeaderFree.vue';
import HeaderBusy from '@/components/HeaderBusy.vue';
import TimelineMini from '@/components/TimelineMini.vue';

export default {
    name: 'HomeView',
    components: {
        HeaderFree,
        HeaderBusy,
        TimelineMini,
    },
};
</script>

<style lang="scss" scoped>
.container {
    margin: 0 auto;
    padding: 32px 24px 24px;
    min-width: 910px;
    color: white;
}
</style>
