Vue.component('ships-view', {
    props: ['ships'],
    data: function () {
        return {}
    },
    methods: {
        getShipImage(ship) {
            var score = 0;
            score += ship.Cannons;
            score += ship.MaxSpeed;
            score += ship.Sight;

            if (score <= 3) {
                return 'assets/ships/Ship_1/Ship_10001.png';
            }
            if (score <= 6) {
                return 'assets/ships/Ship_2/Ship_20001.png';
            }
            if (score <= 9) {
                return 'assets/ships/Ship_3/Ship_30001.png';
            }
            if (score <= 12) {
                return 'assets/ships/Ship_4/Ship_40001.png';
            }
            if (score <= 15) {
                return 'assets/ships/Ship_5/Ship_50001.png';
            }
            if (score <= 18) {
                return 'assets/ships/Ship_6/Ship_60001.png';
            }
            if (score <= 21) {
                return 'assets/ships/Ship_7/Ship_70001.png';
            }
            if (score > 24) {
                return 'assets/ships/Ship_8/Ship_80001.png';
            }
        },
    },
    template: '<div class="ship" v-if="ships.length > 0" v-bind:style="{ backgroundImage: \'url(\' + getShipImage(ships[0]) + \')\' }"></div>'
});
