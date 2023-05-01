<script lang="ts">
    import Radio from "$lib/components/RadioButton.svelte";
    import {onMount} from "svelte";
    import {sendGet} from "$lib/helpers/sender.helper";
    import {VoteModel} from "$lib/models/classes/vote.model";
    import type {CountryModel} from "$lib/models/classes/country.model";
    import {voteEP} from "$lib/models/enums/endpoints.enum";

    export let country:CountryModel
    let vote =  new VoteModel()

    onMount(() => {
        sendGet<VoteModel>(voteEP.BY_COUNTRY + country.slug).then(val => {
            vote = val.body
        })
    })

    const options = [{
        value: 1,
        label: '1'
    },{
        value: 2,
        label: '2'
    },{
        value: 3,
        label: '3'
    },{
        value: 4,
        label: '4'
    },{
        value: 5,
        label: '5'
    },{
        value: 6,
        label: '6'
    },{
        value: 7,
        label: '7'
    },{
        value: 8,
        label: '8'
    },{
        value: 9,
        label: '9'
    },{
        value: 10,
        label: '10'
    }]
</script>

<Radio {options} userSelected={vote.song} legend='Rate the quality of the song'/>
<Radio {options} userSelected={vote.performance} legend='Rate the quality of the performance'/>
<Radio {options} userSelected={vote.costume} legend='Rate the quality of the costumes'/>
<Radio {options} userSelected={vote.props} legend='Rate the quality of the props'/>