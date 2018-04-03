var conf = {
    gdoc: {
        auth : {
            client_id:"896310486270-bhk1b03j6gq8ehut9vib226jjgqdolmi.apps.googleusercontent.com",
            client_secret:"CWGqlpTe8EeJJ7yzqs-HKw00",
            redirect_uri:"urn:ietf:wg:oauth:2.0:oob",
            access_token:"4/AAA5p_LJNFLhLUXdqbFcDe3W769JZG6EaM8LodT8fb4zK-xBMCoy0VU",
            refresh_token:"1/tf7IN8E_iu3-VYE7gEw8l2KyAA5aYT1TKWXmeTL0hkqZRekZ74qaoxzW0SoHCN0_",
        },

        spreadsheet_key:[
            "1YxvcPYt_woGgdYOyYgfIDhy07i_eqBmtzstH75Srbh8"
        ],
    },

    output: {
        path:"../build/",
        data:[
            {type:'tsv'},
            // {type:'csv'},
            // {type:'json'},
        ],
        code:[
            {type:'go', name:'Table_$name', name_kind:'upper'},
            // {type:'cs', name:'Table_$name', name_kind:'upper'}
        ]
    },
}

module.exports = conf