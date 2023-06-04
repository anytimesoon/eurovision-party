<script lang="ts">
    import Spinner from "$lib/components/Spinner.svelte";
    import {currentUser} from "$lib/stores/user.store";
    import type { PageData } from './$types';
    import {browser} from "$app/environment";
    import {countryStore} from "$lib/stores/country.store";
    import {authLvl} from "$lib/models/enums/authLvl.enum";
    import {goto} from "$app/navigation";

    export let data: PageData;

    function setData(data){
        if (browser && data.currentUser != null) {
            $currentUser = data.currentUser
            if ($currentUser.authLvl === authLvl.ADMIN) {
                goto("/admin/countries")
            } else {
                goto("/")
            }
        }
    }

    $: setData(data)
</script>

<Spinner />