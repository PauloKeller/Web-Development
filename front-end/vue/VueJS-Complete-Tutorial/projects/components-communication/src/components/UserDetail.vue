<template>
    <div class="component">
        <h3>You may view the User Details here</h3>
        <p>Many Details</p>
        <!-- Getting data from outside of the component -->
        <p>User Name: {{ switchName() }}</p>
        <p>User Age: {{ userAge }}</p>
        <button @click="resetName">Reset Name</button>
        <button @click="resetFn()">Reset Name</button>
    </div>
</template>

<script>
    import { eventBus } from '../main';

    export default {
        // From outside of the component
        props: {
            myName: {
                // type: [String, Array],
                type: String,
                required: true,
            },
            // call back function from parent
            resetFn: Function,
            userAge: Number
            // user: {
            //     type: Object,
            //     default: function() {
            //         return {
            //             name: 'Paulo',
            //             lastName: 'Keller'
            //         }
            //     }
            // }
        },
        methods: {
            switchName() {
                return this.myName.split("").reverse().join("");
            },
            resetName() {
                this.myName = 'Paulo';
                // Emit event to the parent component
                this.$emit('nameWasReset', this.myName);
            }
        },
        created() {
            // manage the event emited by a child
            eventBus.$on('ageWasEdited', (age) => {
                this.userAge = age;
            }); 
        },
    }
</script>

<style scoped>
    div {
        background-color: lightcoral;
    }
</style>
