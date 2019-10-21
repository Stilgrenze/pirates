Vue.component('events-view', {
    props: ['events'],
    data: function () {
        return {}
    },
    methods: {
        getEventIcon(event) {
            if (event.Message.match(/\[NEW-SHIP\]/)) {
                return 'assets/icons/3.png';
            }
            if (event.Message.match(/\[SHIP-ATTACK\]/)) {
                return 'assets/icons/5.png';
            }
            if (event.Message.match(/\[SHIP-DESTROYED\]/)) {
                return 'assets/icons/7.png';
            }
            if (event.Message.match(/\[PORT-ATTACK\]/)) {
                return 'assets/icons/2.png';
            }
            if (event.Message.match(/\[PORT-LOOTED\]/)) {
                return 'assets/icons/4.png';
            }
        },
    },
    template: `
        <ul class="events list-group">
            <li class="event list-group-item" v-for="(event, index) in events">
                <span class="event-icon" v-bind:style="{ backgroundImage: 'url(' + getEventIcon(event) + ')' }"></span>
                {{ event.Message }}
            </li>
        </ul>
`,
});
