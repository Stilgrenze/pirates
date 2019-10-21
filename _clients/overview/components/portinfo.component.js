Vue.component('port', {
    props: ['port'],
    data: function () {
        return {}
    },
    methods: {
        getOuttimeImage(port) {
            if (port.Outtime > 0) {
                return 'assets/explosion2/ex3.png';
            }
            return 'assets/transparent.png';
        },
    },
    template: `
        <div class="port" v-if="port">
            <div class="explosion" v-bind:style="{ backgroundImage: 'url(' + getOuttimeImage(port) + ')' }"></div>
            <span class="porttext">{{ port.Name }}</span>
        </div>
    `
});
