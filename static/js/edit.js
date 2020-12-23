$('document').ready(function(){
    $('.table #editButton').on('click',function(event){
        event.preventDefault();
        var href=$(this).attr('href');

        $.get(href, function(user, status){
            $('#userFirstNameEdit').val(user.FirstName);
            $('#userLastNameEdit').val(user.LastName);
            $('#userEmailEdit').val(user.Email);
            $('#userMobileNumberEdit').val(user.MobileNumber);
            $('#userAddressEdit').val(user.Address);
            $('#userCityEdit').val(user.City);
            $('#userRoleEdit').val(user.Role);
            $('#userStatusEdit').val(user.Status);
            var editAction = $('#editModal #userUpdateForm').attr('action');
            $('#editModal #userUpdateForm').attr('action', editAction + user.UserId);
        });

        $('#editModal').modal();
    });

    $('table #deleteButton').on('click',function(event){
        event.preventDefault();
        var href=$(this).attr('href');
        $('#deleteModal #userDeleteForm').attr('action', href);
        $('#deleteModal').modal();
    });
});