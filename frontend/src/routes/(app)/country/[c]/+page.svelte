<script lang="ts">
    import Radio from "$lib/components/RadioButton.svelte";

    /** @type {import('../../../index').PageData} */
    export let data;

    // onMount(() => {
    //
    //     const oldCountry = selectedCountry;
    //     sendGet<CountryModel>(countryEP.FIND_ONE + slug).then(val => selectedCountry = val.body);
    //
    //     if(selectedCountry !== oldCountry) {
    //         let i = $partCountryStore.indexOf(oldCountry)
    //         $partCountryStore[i] = selectedCountry
    //     }
    // })

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
{data.slug}
<h1>{data.c.flag} {data.c.name}</h1>

<!--<h3>Song</h3>-->
<!--<p>{selectedCountry.songName} by {selectedCountry.bandName}</p>-->

<h3>Vote!</h3>
<Radio {options} fontSize={16} legend='Score the quality of the song'/>
<Radio {options} fontSize={16} legend='Score the quality of the performance'/>
<Radio {options} fontSize={16} legend='Score the quality of the costumes'/>
<Radio {options} fontSize={16} legend='Score the quality of the props'/>
