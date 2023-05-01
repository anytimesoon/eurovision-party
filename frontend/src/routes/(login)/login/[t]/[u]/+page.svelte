<script>

    import {onMount} from "svelte";
    import {LoginModel} from "$lib/models/classes/login.model";
    import {loginAndGetUsers} from "$lib/helpers/login.helper";
    import {goto} from "$app/navigation";
    import {currentUserStore} from "$lib/stores/user.store";
    import {authLvl} from "$lib/models/enums/authLvl.enum";
    import Spinner from "$lib/components/Spinner.svelte";

    export let data;


    onMount(() => {
        const payload = new LoginModel(data.token, data.userId)

        loginAndGetUsers(payload).then( _ => {
            if ($currentUserStore.authLvl === authLvl.ADMIN) {
                goto("/admin/countries")
            } else if ($currentUserStore.authLvl === authLvl.USER) {
                goto("/")
            } else {
                alert("Something went very wrong")
            }
        })
    })
</script>

<Spinner />