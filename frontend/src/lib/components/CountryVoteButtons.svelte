<script lang="ts">
    import Radio from "$lib/components/RadioButton.svelte";
    import {onMount} from "svelte";
    import {sendCreateOrUpdate, sendGet} from "$lib/helpers/sender.helper";
    import {VoteModel} from "$lib/models/classes/vote.model";
    import type {CountryModel} from "$lib/models/classes/country.model";
    import {voteEP} from "$lib/models/enums/endpoints.enum";
    import {VoteSingleModel} from "$lib/models/classes/voteSingle.model";
    import {currentUserStore} from "$lib/stores/user.store";
    import {catogories} from "$lib/models/enums/categories.enum";

    export let country:CountryModel
    let vote = new VoteModel()

    onMount(() => {
        sendGet<VoteModel>(voteEP.BY_COUNTRY + country.slug).then( val => vote = val.body )
    })

    const options:Array<{ label: string; value: number }>  = [{
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

    function sendVote(n:number, name:String){
        let newVote = new VoteSingleModel($currentUserStore.id, country.slug, name, n)
        sendCreateOrUpdate<VoteSingleModel, VoteModel>(voteEP.UPDATE, newVote).then( val => vote = val.body)
    }
</script>


    <Radio {options} sendVote={sendVote} userSelected={vote.song} name={catogories.SONG} legend='Rate the quality of the song'/>
    <Radio {options} sendVote={sendVote} userSelected={vote.performance} name={catogories.PERFORMANCE} legend='Rate the quality of the performance'/>
    <Radio {options} sendVote={sendVote} userSelected={vote.costume} name={catogories.COSTUME} legend='Rate the quality of the costumes'/>
    <Radio {options} sendVote={sendVote} userSelected={vote.props} name={catogories.PROPS} legend='Rate the quality of the props'/>
