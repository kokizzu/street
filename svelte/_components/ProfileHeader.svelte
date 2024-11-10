<script>
    import {T, isSideMenuOpen, langOptions} from './uiState.js';
    import { Icon } from '../node_modules/svelte-icons-pack/dist';
    import {
			FaSolidBars, FaSolidPowerOff, FaCommentDots
    } from '../node_modules/svelte-icons-pack/dist/fa';
    import {onMount} from 'svelte';
    import {UserUpdateProfile, UserLogout, UserSendFeedback} from '../jsApi.GEN.js';
    import {notifier} from './notifier.js';

    export let user = null;
    export let access = {
        'admin': false,
        'buyer': false,
        'realtor': false,
        'user': false,
    };

    function openSideMenu() {
        isSideMenuOpen.set(!$isSideMenuOpen);
    }

    let selectedLanguage = '';
    onMount(() => {
        console.log('onMount.ProfileHeader =', user)
        selectedLanguage = T.currentLang || 'EN';
    });

    async function updateLang() {
        if (user !== null && user.language !== selectedLanguage) {
            user.language = selectedLanguage;
            await UserUpdateProfile(user, function(res) {
                if (res.error) {
                    notifier.showError(res.error);
                    return;
                }
                user = res.user;
            });
        }
    }

    $: {
        T.changeLanguage(selectedLanguage, async () => await updateLang());
    }

    let showLogoutMenu = false;
    async function userLogout() {
        await UserLogout( {}, function( o ) {
            console.log( o );
            if( o.error ) {
                notifier.showError( o.error );
                return;
            }
            notifier.showSuccess( 'Logged out' );
            window.location = '/';
        } );
    }

    let userMessage = '', showFeedbackDialog = false;
    function sendFeedback() {
        if (userMessage === '') {
            notifier.showError( 'Please enter your feedback');
            return
        }
        UserSendFeedback({userMessage}, function( res ){
            if(res.error) {
                userMessage = '';
                showFeedbackDialog = false;
                notifier.showError(res.error);
                return
            }
            userMessage = '';
            notifier.showSuccess('feedback sent')
            showFeedbackDialog = false;
        })
    }
</script>

{#if showFeedbackDialog}
<div class="backdrop">
    <div class="feedback_dialog">
        <h3>Feedback</h3>
        <textarea bind:value={userMessage} name="feedback" id="feedback" cols="30" rows="10" placeholder="your feedback here"></textarea>
        <div class="buttons">
            <button class="cancel_btn" on:click={() => showFeedbackDialog = false}>Cancel</button>
            <button class="submit_btn" on:click={sendFeedback}>Submit</button>
        </div>
    </div>
</div>
{/if}
<header class='profile_header'>
    <nav class='navbar'>
        <div class='label_menu'>
            {#if access.admin}
                <button on:click|preventDefault={openSideMenu}>
                    <Icon color='#FFF' size={20} src={FaSolidBars}/>
                </button>
            {/if}
            {#if !access.admin}
                <button on:click|preventDefault={() => showLogoutMenu = !showLogoutMenu}>
                    <Icon color='#FFF' size={20} src={FaSolidPowerOff}/>
                </button>
            {/if}
            <p>HapSTR</p>
            {#if showLogoutMenu}
            <div class="logout_menu">
                <button on:click|preventDefault={userLogout}>
                    Logout
                </button>
                <button on:click|preventDefault={() => showLogoutMenu = !showLogoutMenu}>
                    Cancel
                </button>
            </div>
            {/if}
        </div>
        <div class='right_nav'>
            <button class="feedback_btn" on:click={() => showFeedbackDialog = true}>
                <Icon className="feedback_icon" color='#475569' size={16} src={FaCommentDots}/>
                <span>Feedback</span>
            </button>
            <select bind:value={selectedLanguage} id='lang' name='lang'>
                {#each Object.values(langOptions) as lang}
                    <option value={lang}>{lang}</option>
                {/each}
            </select>
<!--            <button class='profile_button' on:click={userLogout} title='Click to logout'>-->
<!--                <img alt='profile' src='/assets/img/team-1-200x200.jpg'/>-->
<!--            </button>-->
        </div>
    </nav>
</header>

<style>

    .backdrop {
        position: fixed;
        z-index: 40;
        background: rgba(41, 41, 41, 0.5);
        width: 100%;
        left: 0;
        top: 0;
        right: 0;
        bottom: 0;
        height: auto;
        display: flex;
        justify-content: center;
        align-items: center;
    }

    .feedback_dialog {
        padding: 20px;
        border-radius: 8px;
        background-color: #FFF;
        display: flex;
        flex-direction: column;
        gap: 20px;
        color: #475569;
        width: 500px;
    }

    .feedback_dialog h3 {
        margin: 0;
        text-align: center;
        font-size: 18px;
    }

    .feedback_dialog textarea {
        border           : 1px solid #CBD5E1;
        background-color : #F1F5F9;
        border-radius    : 8px;
        padding          : 12px;
        resize: none;
    }

    .feedback_dialog textarea:focus {
        border-color : #3B82F6;
        outline      : 1px solid #3B82F6;
    }

    .feedback_dialog .buttons {
        display: flex;
        flex-direction: row;
        align-items: stretch;
        font-weight: 500;
        width: 100%;
        gap: 20px;
    }
    .feedback_dialog .buttons button {
        border-radius: 8px;
        border: none;
        padding: 10px;
        cursor: pointer;
        width: 100%;
        text-transform: capitalize;
    }
    .feedback_dialog .buttons .cancel_btn {
        border: 1px solid #cbd5e1;
        background-color: #f1f5f9;
        color: #f97316;
    }
    .feedback_dialog .buttons .cancel_btn:hover {
        border: 1px solid #f97316;
    }

    .feedback_dialog .buttons .submit_btn {
        background-color : #F97316;
        color            : white;
    }

    .feedback_dialog .buttons .submit_btn:hover {
        background-color : #F58433;
    }

    /* +=========== ==============+ */

    .profile_header {
        background-color : #EF4444;
        height           : fit-content;
        position         : relative;
        padding          : 18px 50px 80px 50px;
        display          : flex;
        flex-direction   : column;
    }

    .profile_header .navbar {
        display         : flex;
        flex-direction  : row;
        justify-content : space-between;
        align-items     : center;
    }

    .profile_header .navbar .label_menu {
        display        : flex;
        flex-direction : row;
        align-items    : center;
        color          : white;
        position: relative;
    }

    .profile_header .navbar .label_menu .logout_menu {
        display: flex;
        flex-direction: column;
        border-radius: 8px;
        filter        : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
        
        background-color: #FFF;
        color: #475569;
        position: absolute;
        z-index: 800;
        border: 1px solid #CBD5E1;
        height: fit-content;
        width: fit-content;
        top: 50px;
    }

    .profile_header .navbar .label_menu .logout_menu button {
        border: none;
        background-color: transparent;
        padding: 0;
        color: #475569;
        font-weight: 600;
        padding: 9px 17px;
    }

    .profile_header .navbar .label_menu .logout_menu button:nth-child(1) {
        color: #EF4444;
    }

    .profile_header .navbar .label_menu .logout_menu button:hover {
        background-color: #F1F5F9;
    }

    .profile_header .navbar .label_menu button {
        padding       : 10px;
        border        : none;
        background    : none;
        border-radius : 5px;
        font-size     : 14px;
        color         : white;
        cursor        : pointer;
    }

    .profile_header .navbar .label_menu button:hover {
        background-color : rgba(255, 255, 255, 0.3);
    }

    .profile_header .navbar .label_menu p {
        font-size   : 17px;
        font-weight : 600;
        line-height : 1.5rem;
        padding     : 0;
        margin      : 0 0 0 15px;
    }

    .profile_header .navbar .right_nav {
        display        : flex;
        flex-direction : row;
        gap            : 20px;
        align-items    : stretch;
    }

    .profile_header .navbar .right_nav .feedback_btn {
        display        : flex;
        flex-direction: row;
        gap: 5px;
        align-items    : center;
        background-color: #FFF;
        border-radius: 5px;
        padding       : 6px 8px;
        border: none;
        color: #475569;
        font-weight: 600;
        cursor: pointer;
        filter        : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
    }

    .profile_header .navbar #lang {
        padding       : 6px 8px;
        width         : 60px;
        border        : none;
        border-radius : 5px;
        color         : #1080E8;
        font-weight   : 600;
        cursor        : pointer;
        background-color: #FFF;
        filter        : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));
    }

    .profile_header .navbar #lang:hover,
    .profile_header .navbar .right_nav .feedback_btn:hover {
        background-color : #F1F5F9;
    }

    .profile_header .navbar #lang:focus {
        outline : none;
    }

    /*.profile_header .navbar .right_nav .profile_button {*/
    /*    padding       : 0;*/
    /*    border        : none;*/
    /*    margin        : 0;*/
    /*    width         : fit-content;*/
    /*    height        : fit-content;*/
    /*    border-radius : 50%;*/
    /*    filter        : drop-shadow(0 10px 8px rgb(0 0 0 / 0.04)) drop-shadow(0 4px 3px rgb(0 0 0 / 0.1));*/
    /*}*/
    
    /*.profile_header .navbar .right_nav .profile_button:hover {*/
    /*    box-shadow : 0 0 0 3px rgba(255, 255, 255, 0.5);*/
    /*    cursor     : pointer;*/
    /*}*/
    
    /*.profile_header .navbar .right_nav .profile_button > img {*/
    /*    width         : 50px;*/
    /*    height        : 50px;*/
    /*    border-radius : 9999px;*/
    /*}*/

    /* Responsive to mobile device */
    @media only screen and (max-width : 768px) {

        .feedback_dialog {
            width: 100%;
            margin: auto 20px;
        }

        .profile_header {
            padding : 15px 20px 80px 20px !important;
        }

        .profile_header .navbar .label_menu button {
            padding : 0;
        }

        .profile_header .navbar .right_nav {
            gap : 10px;
        }

        .profile_header .navbar #lang {
            padding   : 3px 5px;
            font-size : 12px;
            width     : 50px;
        }

        .profile_header .navbar .feedback_btn {
            padding   : 3px 5px;
            font-size : 12px;
        }

        :global(.feedback_icon) {
            width: 12px;
            height: 12px;
        }

        /*.profile_header .navbar .right_nav .profile_button > img {*/
        /*    width  : 40px;*/
        /*    height : 40px;*/
        /*}*/

        .profile_header .navbar .label_menu .logout_menu {
            top: 30px;
        }
    }
</style>
