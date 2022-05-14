## Show avaliable participants

# Description
Helping in select participants for interview according rules like "not more 2 for interviewer per week" and show the avaliable slots

# PG
test test

# API
/show?startDay=&endDay=

{
    {day}: {
        {id}: {
            name: "",
            slots: {
                {priorityLevel}: [
                    {
                        from: {time},
                        to: {time},
                    },

                ]
            }
        },
        ...
    },
    ...
}

/select?id=&from=&to= // check avaliability without priority

{
    event_id: {id}
    error: {reason}
}

/cancel?event_id=