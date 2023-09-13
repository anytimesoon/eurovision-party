<script lang="ts" xmlns="http://www.w3.org/1999/html">
    import {CountryModel} from "$lib/models/classes/country.model";
    import { enhance } from '$app/forms';

    let form;
    export let country:CountryModel = new CountryModel("", "", "", "", "", false)

</script>

<form method="POST" action="?/update" use:enhance bind:this={form}>
    <input name="slug" type="hidden" bind:value={country.slug}>
    <input type="checkbox" name="participating" class="hidden" bind:checked={country.participating}>

    <div>
        <label for="{country.slug}-song">Song Title</label>
        <input id="{country.slug}-song" name="songName" type="text" bind:value={country.songName} >
    </div>

    <div>
        <label for="{country.slug}-act">Act Name</label>
        <input id="{country.slug}-act" name="bandName" type="text" bind:value={country.bandName} >
    </div>


    <button>Save</button>
</form>