import 'user_info_card.dart';
import 'user_info_card.dart';
import 'user_info_card.dart';
import 'user_info_card.dart';
class GroupBan{
	String groupId;
	UserInfoCard lastModifiedBy;
	UserInfoCard createdBy;
	String dateBanned;
	String dateExpires;
	String comment;
	UserInfoCard bungieNetUserInfo;
	UserInfoCard destinyUserInfo;
	GroupBan(
		String this.groupId,
		UserInfoCard this.lastModifiedBy,
		UserInfoCard this.createdBy,
		String this.dateBanned,
		String this.dateExpires,
		String this.comment,
		UserInfoCard this.bungieNetUserInfo,
		UserInfoCard this.destinyUserInfo,
	);

	static GroupBan fromMap(Map<String, dynamic> data){
		if(data == null) {
			return null;
		};
		return new GroupBan(
				data['groupId'],
				UserInfoCard.fromMap(data['lastModifiedBy']),
				UserInfoCard.fromMap(data['createdBy']),
				data['dateBanned'],
				data['dateExpires'],
				data['comment'],
				UserInfoCard.fromMap(data['bungieNetUserInfo']),
				UserInfoCard.fromMap(data['destinyUserInfo']),
		);
	}

	static List<GroupBan> fromList(List<dynamic> data){
		if(data == null) {
			return null;
		};
		List<GroupBan> list = new List();
    data.forEach((item) {
      list.add(GroupBan.fromMap(item));
    });
    return list;
	}

	Map<String, dynamic> toMap(){
		Map<String, dynamic> data = new Map();
			data['groupId'] = this.groupId;
			data['lastModifiedBy'] = this.lastModifiedBy.toMap();
			data['createdBy'] = this.createdBy.toMap();
			data['dateBanned'] = this.dateBanned;
			data['dateExpires'] = this.dateExpires;
			data['comment'] = this.comment;
			data['bungieNetUserInfo'] = this.bungieNetUserInfo.toMap();
			data['destinyUserInfo'] = this.destinyUserInfo.toMap();
		return data;
	}
}