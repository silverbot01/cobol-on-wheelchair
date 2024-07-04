

        identification division.
        program-id. testpage.

        data division.
        working-storage section.

        01 the-vars.

            03 COW-vars OCCURS 99 times.

                05 COW-varname      pic x(99).
                05 COW-varvalue     pic x(99).

        procedure division.

            MOVE "testvalue" to COW-varname(1)
            MOVE "successful" to COW-varvalue(1)
            call 'cowtemplate' using the-vars "test.cow"


        goback.

        end program testpage.