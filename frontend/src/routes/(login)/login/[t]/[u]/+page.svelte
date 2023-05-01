<style>
    .spinner {
        position: fixed;
        z-index: 1;
        left: 0;
        right: 0;
        top: 0;
        bottom: 0;
        background: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' width='40' height='40' viewBox='0 0 50 50'%3E%3Cpath d='M28.43 6.378C18.27 4.586 8.58 11.37 6.788 21.533c-1.791 10.161 4.994 19.851 15.155 21.643l.707-4.006C14.7 37.768 9.392 30.189 10.794 22.24c1.401-7.95 8.981-13.258 16.93-11.856l.707-4.006z'%3E%3CanimateTransform attributeType='xml' attributeName='transform' type='rotate' from='0 25 25' to='360 25 25' dur='0.6s' repeatCount='indefinite'/%3E%3C/path%3E%3C/svg%3E") center / 50px no-repeat;
    }
</style>

<script>

    import {onMount} from "svelte";
    import {LoginModel} from "$lib/models/classes/login.model";
    import {loginAndGetUsers} from "$lib/helpers/login.helper";
    import {goto} from "$app/navigation";
    import {currentUserStore, userStore} from "$lib/stores/user.store";
    import {authLvl} from "$lib/models/enums/authLvl.enum";

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

{data.token}
{data.userId}
{$currentUserStore}

<div class="spinner"></div>