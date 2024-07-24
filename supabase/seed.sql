CREATE POLICY "allow all 1ffg0oo_0" ON storage.objects
    FOR SELECT TO public
        USING (bucket_id = 'images');

CREATE POLICY "allow all 1ffg0oo_1" ON storage.objects
    FOR INSERT TO public
        WITH CHECK (bucket_id = 'images');

INSERT INTO setting_types(setting_type)
    VALUES ('image'),
('link'),
('text'),
('richtext'),
('address'),
('phone'),
('number'),
('percent');


INSERT into settings(
  setting_type_id,
  setting_key,
  setting_value
) values 
(3 , 'site_title' , 'بزنس برو | Business Pro'),
(3 , 'copyrights' , ' جميع حقوق النشر محفوظة © بزنس برو. '),
(3 , 'instagram' , 'https://www.instagram.com/'),
(3 , 'email' , 'BusinessPro@gmail.com'),
(3 , 'phone' , '+966 55 165 1833'),
(3 , 'address' , ' السعودية , الرياض , حي النرجس , طريق انس بن مالك '),
(3 , 'twitter' , 'https://www.twitter.com/'),
(3 , 'youtube' , 'https://www.youtube.com/'),
(3 , 'bank_name' , 'البنك الاهلي السعودي'),
(3 , 'bank_iban' , 'SA9010000012472813000102'),
(3 , 'bank_account' , '12472813000102'),
(3 , 'request_success_notification' , 'تم التسجيل بنجاح سنقوم بالتواصل معكم'),

(7 , 'programms_count' , '100'),
(7 , 'beneficiary_count' , '2000'),
(7 , 'projects_count' , '350'),
(7 , 'consultunts_count' , '1100'),
(1 , 'banner_image' , 'initial/global_slider_1.webp'),
(1 , 'about_image' , 'initial/global_about.webp'),

(4 , 'banner_header' , '<h1 class="text-secondary" style="margin-bottom : 20px">
              خدمات استشارية <br /><strong class="mr-110"> لتطوير اعمالك </strong>
            </h1>'),
(4 , 'banner_text' , '<p class="text-black text-h6 w-100 mr-110">
              نساعد الشركات و المنظمات على تحقيق النمو و الاستدامة <br /> لصناعة حلول مؤثرة
            </p>'),
(4 , 'blog_page_description' , 'يسعدنا تواصلك معنا للرد علي استفساراتك او لبحث إمكانية العمل معا ومساعدتك علي تطوير شرتكتك'),

(4 , 'vision' , '<p class="text-h6"> الشركة السعودية الرائدة في الحلول  <br> الإستشارية والإستثمارية</p>'),
(4 , 'about' , '<p class="text-h6 aligned-paragraph">
          نطرح على الطاولة العديد من الخدمات الاستشارية القيمة المختلفة التي
          تؤثر بشكل مباشر على الإدارات الأساسية الشركة السعودية الرائدة في تقديم
          خدمات الاستشارات الإدارية والمالية وحلول الاعمال
        </p>
        <p class="text-h6 aligned-paragraph">
          الشركة السعودية الرائدة في تقديم خدمات الاستشارات الإدارية والمالية
          وحلول الاعمال
        </p>
        <p class="text-h6 aligned-paragraph">
          نطرح على الطاولة العديد من الخدمات الاستشارية القيمة المختلفة التي
          تؤثر بشكل مباشر على الإدارات الأساسية الشركة السعودية الرائدة في تقديم
          خدمات الاستشارات الإدارية والمالية وحلول الاعمال
        </p>'),
(4 , 'message' , '<p class="text-h6">تطوير المنظمات وقادة الاعمال نحو المستقبل <br> بـرؤيـة مـخـتـلـفـة</p>'),
(4 , 'values' , '<p class="text-h6">تطوير الأداء، تحقيق النجاح المستدام، <br>واكتشاف الحلول المبتكرة</p>');
            
INSERT INTO request_statuses(request_status)values('معلق') , ('مؤكد') ,   ('ملغي');
INSERT INTO category_types(category_type)values('projects') ,  ('blog');

INSERT INTO categories(category_name ,category_type_id)values ('الريادة و الابتكار' , 1),('دراسة و تحليل السوق' , 1),('حصص واسهم الشركات' , 1),('نموذج عمل تجاري' , 1),('اقتصاد' , 2),('استشارة' , 2),('اعمال' , 2);

INSERT INTO events(
    event_name,
    event_location,
    event_location_url,
    constructor_name,
    constructor_title,
    constructor_image,
    tags,
    price,
    discount,
    shab_discount,
    event_plan,
    event_goals,
    event_breif,
    event_description,
    event_video,
    event_date,
    event_start_time,
    event_end_time,
     event_hours,
    category_id,
    event_image
) values (
    'دورة تدريبية لبناء خطط الاعمال للشركات',
    'الرياض',
    'https://maps.app.goo.gl/GnNgWKsfSwACRxoi9',
    'أ / أحمد المنهبي ',
    'خبير اقتصادي',
    'initial/events_single3.webp',
    array['الاحدث'],
    575,
    15,
    25,
    array['{
      "title": "محاكاة عملية",
      "breif": "تجربة عملية تفاعلية تمنحك الفرصة للتعلم من خلال محاكاة أوضاع وظروف العمل الحقيقية"
    }'::jsonb,
    '{
      "title": "أدوات قياس",
      "breif": "تعلم استخدام أدوات القياس الفعالة لتحليل أداء الشركة واتخاذ القرارات الاستراتيجية المبنية على البيانات"
    }'::jsonb,
    '{
      "title": "نقاشات نشطة",
      "breif": "المشاركة في نقاشات مفصلة وتبادل الأفكار والآراء مع المشاركين الآخرين لتعزيز التعلم وتوسيع المدارك"
    }'::jsonb,
    '{
      "title": "شبكة علاقات",
      "breif": "بناء علاقات قوية ومفيدة مع المشاركين والمحاضرين والخبراء في مجال الأعمال وتوسيع شبكتك المهنية"
    }'::jsonb,
    '{
      "title": "نمذجة الابتكار بواسطة الذكاء الاصطناعي (AI)",
      "breif": "استخدام تقنيات الذكاء الاصطناعي لتحليل البيانات واستخراج الصيحات وتوجيه عملية اتخاذ القرارات للابتكار في المشاريع"
    }'::jsonb],
    array[
    'توفير مجموعة متنوعة من الدورات التدريبية المكثفة لتطوير مهارات الإدارة والتسويق والتمويل والابتكار في بناء خطط الأعمال الناجحة.',
    'تعزيز قدرات المشاركين على استخدام أدوات قياس الأداء وتحليل البيانات لتحقيق تحسين مستمر واتخاذ القرارات الاستراتيجية المبنية على الحقائق.',
    'تطوير مهارات التواصل والتفاوض والعمل الجماعي من خلال المشاركة في نقاشات نشطة ومحاكاة عملية لتعزيز التعلم وبناء شبكات علاقات مهنية قوية.',
    'تعريف المشاركين بأحدث تقنيات الذكاء الاصطناعي واستخدامها في نمذجة الابتكار وتحليل البيانات لتحقيق تفوق ريادي في مجالات الأعمال المختلفة',
    'تمكين المشاركين من وضع خطط استراتيجية شاملة للمشاريع وتحليل السوق واستراتيجية التسويق وتخطيط الموارد وإدارة المخاطر لتحقيق النجاح المستدام في الأعمال التجارية'
  ],
    ' تركز الورشة على تطوير وتعزيز الإبداع والابتكار الريادي والتفوق بأدوات الذكاء الاصطناعي والعمل على تشجيع التفكير الخلاق من خلال الاستثمار المستقبلي لتحقيق أفكار جديدة وجريئة تتوافق مع تطلعات المستقبل',
   '<p>تركز الورشة على تطوير وتعزيز الإبداع والابتكار الريادي والتفوق بأدوات الذكاء الاصطناعي والعمل على تشجيع التفكير الخلاق من خلال الاستثمار المستقبلي لتحقيق أفكار جديدة وجريئة تتوافق مع تطلعات المستقبل.</p>

<p>إدارة العمليات هو مساق مجاني مقدّم من إدراك يساعد المتعلم على استيعاب أهم المفاهيم والمهام الرئيسية المتعلقة بموضوع إدارة العمليات، ليتمكن من شق طريقه المهني في هذا المجال بنجاح وممارسة مسؤولياته وصلاحياته بالشكل الصحيح وفي الوقت والمكان المناسبين. إضافةً إلى ذلك، يوضح المساق أهمية إدارة العمليات والأجزاء الرئيسية فيها وكيفية تصميم وتطوير المنتجات والخدمات في المؤسسة.</p>

<p>تكمن أهمية إدارة العمليات في تحقيق التكامل بين الأقسام المختلفة للمؤسسة، وتُعَد إدارة العمليات من المواضيع الهامة التي تُعنى بها المؤسسات والشركات الكبرى تدور تفاصيل عملها حول كيفية تقديم الخدمات وإنتاج البضائع والسلع لتشمل عدداً من الجوانب، ابتداءً من تصميم المنتجات وإدارة المستودعات والمخازن وحتى تحسين وتطوير أداء إدارة العمليات في الشركات.</p>

<p>تُمثل إدارة العمليات حلقة الوصل بين جميع الوظائف داخل المؤسسات والشركات الكبرى، ومن الضروري جداً الوعي التام بجميع مهام إدارة العمليات ومسؤوليات مدير العمليات كي تسير الوظائف ب
',
'IG4-i4XTFr0',
'09/09/2024',
'04:00:00',
'09:00:00',
'12',
1,
'initial/events_p1.webp'

),
(
    'صناعة الريادة والابتكار في المنظمات والشركات',
    'الرياض',
    'https://maps.app.goo.gl/GnNgWKsfSwACRxoi9',
    'أ / أحمد المنهبي ',
    'خبير اقتصادي',
    'initial/events_single3.webp',
    array['الاحدث'],
    480,
    15,
    25,

    array['{
      "title": "محاكاة عملية",
      "breif": "تجربة عملية تفاعلية تمنحك الفرصة للتعلم من خلال محاكاة أوضاع وظروف العمل الحقيقية"
    }'::jsonb,
    '{
      "title": "أدوات قياس",
      "breif": "تعلم استخدام أدوات القياس الفعالة لتحليل أداء الشركة واتخاذ القرارات الاستراتيجية المبنية على البيانات"
    }'::jsonb,
    '{
      "title": "نقاشات نشطة",
      "breif": "المشاركة في نقاشات مفصلة وتبادل الأفكار والآراء مع المشاركين الآخرين لتعزيز التعلم وتوسيع المدارك"
    }'::jsonb,
    '{
      "title": "شبكة علاقات",
      "breif": "بناء علاقات قوية ومفيدة مع المشاركين والمحاضرين والخبراء في مجال الأعمال وتوسيع شبكتك المهنية"
    }'::jsonb,
    '{
      "title": "نمذجة الابتكار بواسطة الذكاء الاصطناعي (AI)",
      "breif": "استخدام تقنيات الذكاء الاصطناعي لتحليل البيانات واستخراج الصيحات وتوجيه عملية اتخاذ القرارات للابتكار في المشاريع"
    }'::jsonb],
    array[
    'توفير مجموعة متنوعة من الدورات التدريبية المكثفة لتطوير مهارات الإدارة والتسويق والتمويل والابتكار في بناء خطط الأعمال الناجحة.',
    'تعزيز قدرات المشاركين على استخدام أدوات قياس الأداء وتحليل البيانات لتحقيق تحسين مستمر واتخاذ القرارات الاستراتيجية المبنية على الحقائق.',
    'تطوير مهارات التواصل والتفاوض والعمل الجماعي من خلال المشاركة في نقاشات نشطة ومحاكاة عملية لتعزيز التعلم وبناء شبكات علاقات مهنية قوية.',
    'تعريف المشاركين بأحدث تقنيات الذكاء الاصطناعي واستخدامها في نمذجة الابتكار وتحليل البيانات لتحقيق تفوق ريادي في مجالات الأعمال المختلفة',
    'تمكين المشاركين من وضع خطط استراتيجية شاملة للمشاريع وتحليل السوق واستراتيجية التسويق وتخطيط الموارد وإدارة المخاطر لتحقيق النجاح المستدام في الأعمال التجارية'
  ],
    ' تركز الورشة على تطوير وتعزيز الإبداع والابتكار الريادي والتفوق بأدوات الذكاء الاصطناعي والعمل على تشجيع التفكير الخلاق من خلال الاستثمار المستقبلي لتحقيق أفكار جديدة وجريئة تتوافق مع تطلعات المستقبل',
   '<p>تركز الورشة على تطوير وتعزيز الإبداع والابتكار الريادي والتفوق بأدوات الذكاء الاصطناعي والعمل على تشجيع التفكير الخلاق من خلال الاستثمار المستقبلي لتحقيق أفكار جديدة وجريئة تتوافق مع تطلعات المستقبل.</p>

<p>إدارة العمليات هو مساق مجاني مقدّم من إدراك يساعد المتعلم على استيعاب أهم المفاهيم والمهام الرئيسية المتعلقة بموضوع إدارة العمليات، ليتمكن من شق طريقه المهني في هذا المجال بنجاح وممارسة مسؤولياته وصلاحياته بالشكل الصحيح وفي الوقت والمكان المناسبين. إضافةً إلى ذلك، يوضح المساق أهمية إدارة العمليات والأجزاء الرئيسية فيها وكيفية تصميم وتطوير المنتجات والخدمات في المؤسسة.</p>

<p>تكمن أهمية إدارة العمليات في تحقيق التكامل بين الأقسام المختلفة للمؤسسة، وتُعَد إدارة العمليات من المواضيع الهامة التي تُعنى بها المؤسسات والشركات الكبرى تدور تفاصيل عملها حول كيفية تقديم الخدمات وإنتاج البضائع والسلع لتشمل عدداً من الجوانب، ابتداءً من تصميم المنتجات وإدارة المستودعات والمخازن وحتى تحسين وتطوير أداء إدارة العمليات في الشركات.</p>

<p>تُمثل إدارة العمليات حلقة الوصل بين جميع الوظائف داخل المؤسسات والشركات الكبرى، ومن الضروري جداً الوعي التام بجميع مهام إدارة العمليات ومسؤوليات مدير العمليات كي تسير الوظائف ب
',
'IG4-i4XTFr0',
'09/09/2024',
'04:00:00',
'09:00:00',
'12',
1,
'initial/events_p2.webp'

);


insert into services(
   "service_name",
  service_image,
  tags,
  breif 
) values(
  'الاستشارات الإدارية',
  'initial/service_money.webp',
  array['مميز']::varchar[],
  '<p class="text-h6">نعمل على تحليل البيانات وتقييم الأداء بهدف التخطيط وتطوير الاستراتيجيات لتحسين وزيادة الإنتاجية وتخفيض التكاليف.</p>'
),
(
  'الاستشارات المالية ',
  'initial/service_marketing.webp',
  array['مميز']::varchar[],
  '<p class="text-h6">تحليل وتقييم البيانات المالية وتحسين الأداء المالي للموارد المالية وتقليل الاعباء المالية بشكل فعال لتحقيق الأرباح المستدامة.  </p>'
),
(
  'الحلول الاستثمارية ',
'initial/service_invest.webp',
  array[]::varchar[],
'<p class="text-h6">نقدم حلول وخطط استثمارية متنوعة في كيفية الاستثمار بشكل فعال لتحقيق أفضل عوائد ممكنة مستدامة على المدى الطويل.  </p>'
),
(
 ' الاستشارات التسويقية ',
'initial/service_marketing.webp',
  array[]::varchar[],
'<p class="text-h6">مساعدة الشركات على تحسين استراتيجيات التسويق وتحقيق أهدافها في إيصال منتجاتها أو خدماتها إلى الجمهور المستهدف لتحقيق نمو في المبيعات</p>'
);




insert into programs(
   "program_name",
  program_image,
  breif 
) values(
   'برنامج الشاب الريادي',
   'initial/program_shabr.webp',
  '<p class="text-h6">يساهم البرنامج في تمكين رواد الاعمال من تبادل الخبرات وتنمية المهارات وتطوير واستدامة اعمالهم ضمن مجتمع اعمال تفاعلي افتراضي.</p>'
),
(
   'برنامج التلمذة المهنية',
   'initial/program_talmza.webp',
  '<p class="text-h6">يعزز اكتـسـاب الثقة والوعي للمتدربين من خلال التدريب على رأس العمل، حيث يوفر تجربة مهنية تساهم في تـمـكـيـن الـشـبـاب والـشـابـات مـن مـهـارات سـوق الـعـمـل السعودي بما يتوافق مع رؤية المملكة 2030</p>'
),
(
   'برنامج تطبيقات الاعمال',
 'initial/program_bznsplan.webp',
'<p class="text-h6">برنامج عبارة عن مجموعة من ورش عمل تطبيقية مخصص لتعزيز وتحسين الكفاءة والإنتاجية في الأعمال بشكل محترف لتحقيق التحسين والابتكار في سياق الأعمال.</p>'
),
(
  'برنامج مدراء الاستثمار',
 'initial/program_modara.webp',
'<p class="text-h6">برنامج تدريبي تفاعلي يعزز تعلم فنون ومهارات إدارة المال وضخها في المشاريع الاستثمارية بما ينعكس إيجاباً على الأداء المالي.</p>'
);

insert into testemonials(
  testemonial_name,
  testemonial_title,
  testemonial_image,
  breif
)values ( 'عبدالعزيز السريع',
 'مستثمر ورجل اعمال',
 'initial/testemonial_abdelaziz.webp',
 'المشاريع المطروحة للإستثمار فريدة من نوعها ومشاركتنا معكم أتت من تميز الطرح '),
(
 'نبيل القرعاوي',
 'مستثمر ورجل اعمال',
 'initial/testemonial_nabil.webp',
 'فريق عمل بزنس برو متميز وذو خبرة واسعة ساهم في تقديم أعمال نوعيه لرواد الأعمال '),
(
 'مازن بترجي',
 'مستثمر ورجل اعمال',
 'initial/testemonial_mazen.webp',
 'شاركتنا في الملتقى الإستثماري حفزت لدينا الرغبة في التواصل البناء للإستثمار في الفرص الواعدة '),
(
 'بندر بن دعجم',
 'الرئيس التنفيذي لحديد بن دعجم',
 'initial/testemonial_bandr.webp',
 'تجربتنا مع بزنس برو منحتنا خوض تجربة عملية<br/> ادت الى نتائج مثمرة '),
(
 'جمال الزامل',
 'مستثمر ورجل اعمال',
 'initial/testemonial_gamal.webp',
 'برامج قيادات الأعمال لتأهيل القيادات الواعدة خطوة استراتيجية لمستقبل الشركات العائلية');

insert into team_members(
  team_member_name,
  job_title,
  team_member_image
) values ( 'احمد المنهبي',
 'المدير عام',
 'initial/team_manhaby.webp'),
 (
 'عبدالله حسان',
 'المشاريع',
 'initial/team_abdallah.webp'),
 (
 'رحاب مغربي',
 'العمليات',
 'initial/team_rehab.webp'),
 (
 'رغدة مرغلان',
 'تطوير الاعمال',
 'initial/team_raghda.webp'),
 (
 'احمد درويش',
 'تقنية المعلومات ',
 'initial/team_darwish.webp'),
 (
 'اسامة المنهبي',
 'الخدمات المشتركة ',
 'initial/team_osama.webp'),
 (
 'يوسف نضال',
 'التصميم و الميديا ',
 'initial/team_yossuf.webp');

insert into blogs (
  blog_name , 
  blog_image , 
  tags,
  breif , 
  content , 
  category_id , 
  user_id , 
  views , 
  date_time , 
  links 
) values (
  'كيف تساهم الاستشارات في تحسين شركتك', 
  'initial/blog_1.webp', 
  array['الاحدث' , 'الاكثر قراءة' ]::varchar[],
  'تتجه شركة بيزنيس برو نحو نمو ديناميكي ومستدام مستهدفةازدهار ورفاهية موظفيها والمجتمع والمساهمة في اقتصاد قطر المتنامي أن تنتهي صلاحية تأشيرة الدخول الخاصة بك',
   '<p> إليك المثال العملي التالي الذي يشرح ماهية البرمجة:
</p>
<p>إن كنت تتوقع زيارة صديق لك اليوم، واتصل بك ليقول لك: "أنا واقف بجانب الحديقة ولا أعرف كيف أصل إلى منزلك". أنت عادةً تمر كل يوم من جانب الحديقة وتعرف الطريق بينها وبين منزلك شبرًا بشبر. برأيك هل ينفع إن قلت له: "منزلي معروف وقريب من الحديقة وأنا كل يوم أمر من جانبها"؟ لا، بالتأكيد. تحتاج إلى أن تقسِّم المشكلة إلى أجزاء تمثل خطوات بسيطة يستطيع صديقك فهمها وتنفيذها.
</p>
<p>مثلًا، أخبره أن ينفذ الأوامر التالية: "سر إلى الأمام عشرة أمتار" ثم "اتجه إلى اليمين" ثم "سر إلى نهاية الشارع" ثم "اتجه إلى اليسار". أخبره بعد ذلك: "عُدَّ الأبنية الموجودة على اليسار حتى تصل إلى البناء الرابع" ثم "اصعد إلى الطابق الثاني" ثم "اطرق على الباب الذي سيظهر أمامك". مبارك! بهذه الطريقة، تستطيع أن تدل صديقك على منزلك بدقة. البرمجة هي الشيء نفسه تمامًا. فهل ترى التعابير المكتوبة بين قوسين؟ إنها التعابير التي تكتب بإحدى لغات البرمجة والتي تخاطب الحاسوب بدلًا من صديقك السابق.
</p>
<p>لغات البرمجة هي مجموعة من المفردات والقواعد اللغوية التي تشكل لغةً وسيطةً للتخاطب مع الحاسوب وأمره بتنفيذ تعليمات وأشياء محدَّدة. فلا الحاسوب يفهم لغة البشر ولا البشر يفهمون لغة الحاسوب، لذا كان هنالك حاجة ملحة لوجود لغة وسيطة يفهمها كلاهما؛ نتيجةً لذلك، انبثق مفهوم لغة البرمجة.
</p>
<p>بعبارة أخرى، لو أردنا أن نقول للحاسوب "افعل كذا"، فسنحتاج إلى لغةٍ مشتركةٍ بيننا وبينه ليفهم ما نبتغيه، وهنا يأتي دور لغات البرمجة، إذ يمكنك أن تعدّ لغات البرمجة على أنها وسيط بين المبرمج والحاسوب.
</p>
<p>يهتم المبرمج بالتفكير في تسلسل الخطوات التي على الحاسوب القيام بها لإتمام العمل المطلوب منه (مثل حساب العمر اعتمادًا على تاريخ الولادة)، ثم كتابة هذه الخطوات بترتيب منطقي بإحدى لغات البرمجة.
</p>
<p>ربما لاحظتَ في الجملة السابقة أن جزءًا من مهمة المبرمج هو التفكير المنطقي، وهذا يجعلنا ننتقل إلى السؤال الشائع "هل أستطيع تعلم البرمجة وأصبح مبرمجًا؟" أو "هل أنا مؤهل لأصبح مبرمجًا؟".
    لماذا تتعلم البرمجة؟</p>
<p>يبدو أن تعلم البرمجة ليس بالصعوبة التي توقعتها، لكنك تريد حافزًا يجعلك تتعلم البرمجة.
</p>
<p>تسمع كثيرًا أن البرمجة هي مجال المستقبل، وأن وظائف المبرمجين ستكتسح مجال التوظيف في السنوات القادمة؟ أستطيع أن أؤكد لك ذلك، كما أنَّ وظائف البرمجة هي من أعلى الوظائف دخلًا.
</p>
<p>فلو كنت تريد بدء مشوارك الاحترافي وتريد عملًا مستقرًا وذا دخلٍ ممتاز، فإن تعلم البرمجة والعمل بها هو أفضل خيارٍ أمامك.
</p>
<p>وظائف البرمجة مريحة عمومًا، فالعمل كله مكتبي أمام حاسوب في بيئة مريحة ومناسبة، وأغلبية الشركات تتبع نظام العمل 40 ساعة في الأسبوع (أي 5 أيام لمدة 8 ساعات يوميًا)، ولا تغفل عن قدرتك على العمل عن بعد من خلال الانترنت أو كمستقل في أوقات فراغك.
</p>
<p>تعلم البرمجة سيوسع أفق تفكيرك كثيرًا، خصوصًا أن تعاملك مع الحاسوب يتبع إلى التفكير المنطقي، وستجد أن البرمجة ستسهل لك القيام بأمور أخرى في الحاسوب.
    ما عليك معرفته لتصبح مبرمجًا
    </p>
<p>يتردد الكثيرون في تعلم البرمجة متذرعين بأن مستواهم في الرياضيات ليس ممتازًا، وهذا ليس صحيحًا، فصحيحٌ أنَّ هنالك أمور تعترضك أثناء أداء عملك كمبرمج تتطلب خبرة في الرياضيات، إلا أنَّه قد تمر عليك فترات طويلة لا تحتاج فيها إلى مسائل رياضية.
</p>
<p>
    كل ما يلزمك للبدء في تعلم البرمجة هو الأساسيات التي يعرفها الجميع. إلى حين اعتراضك أية مسألة أو مشكلة تتطلب مهارة في الرياضيات، هنالك الكثير من المصادر والمراجع التي تستطيع الرجوع إليها آنذاك. بعبارة أخرى، أجِّل هذا الأمر قليلًا ولا تخف. الأهم من ذلك هو أن تكون قادرًا على التفكير بشكل منطقي.
    التفكير المنطقي
    </p>
<p>التفكير المنطقي هو المهارة التي تجمع كافة المبرمجين تحت مظلة واحدة، وهي أساس كتابة الخوارزميات، إذ يجب أن تكون قادرًا على اكتساب هذه المهارة وتطويرها.
    الخوارزميات
    </p>
<p>كلمة "الخوارزميات" هي الكلمة المرعبة التي ينفر منها البعض، فكل ما يتخيلونه عند ذكرها هو الرياضيات المعقدة والمعادلات الطويلة والرموز العجيبة، لكن الأمر بسيط جدًا؛ فالخوازرميات هي تطبيقٌ للتفكير المنطقي في خطوات متسلسلة واضحة تمامًا لحل مشكلة ما.
</p>
<p>لكي أوضِّح لك أن الخوارزميات ليست أمرًا معقدًا، سأخبرك بكيفية كتابة برنامج يسأل المستخدم عن سنة ميلاده، ثم يعيد عمره الحالي بالسنوات. </p>
',
1,
1,
1,
'2024-01-21 15:48:48',
'[{"name": "Facebook", "url": "https://facebook.com/example"},{"name": "Twitter", "url": "https://twitter.com/example"}]'
) ,  
(
  'الإستراتيجية الوطنية للتقنية الحيوية', 
  'initial/blog_2.webp', 
  array[ 'الاكثر قراءة' ]::varchar[],
  'تعزيزًا لمكانة المملكة كدولة رائدة في قطاع التقنية الحيوية، تجسد الاستراتيجية الوطنية للتقنية الحيوية ذلك، وتتسع آفاقها لتعبّر عن بداية رحلةٍ تحوليةٍ للمملكة وللعالم. ',
  'تعزيزًا لمكانة المملكة كدولة رائدة في قطاع التقنية الحيوية، تجسد الاستراتيجية الوطنية للتقنية الحيوية ذلك، وتتسع آفاقها لتعبّر عن بداية رحلةٍ تحوليةٍ للمملكة وللعالم. 


تستهدف الاستراتيجية أن تصبح المملكة تجمعاً عالمياً رائداً في مجال التقنية الحيوية، إلى جانب تحقيق مستوى عالٍ من الاكتفاء الذاتي وإحداث أثر اجتماعي واقتصادي إيجابي، وذلك من خلال تركيز مجهوداتها على أربع توجهات استراتيجية ذات أولوية.

    اللقاحات:

    تُعزز التقنية من إمكانيات التصنيع المتكاملة للقاحات، التي تسهم في تحقيق مستويات عالية من الاكتفاء الذاتي، إلى جانب تطوير فرص تفتح مجالًا أوسع للتصدير إلى منطقة الشرق الأوسط وشمال أفريقيا، وتمكين جهود البحث والتطوير باستخدام أحدث التقنيات المبتكرة؛ لتُصبح المملكة مركزًا رائدًا في مجال التطوير النهائي للقاحات في المنطقة.

    التصنيع الحيوي والتوطين:

    تُمكّن التقنية من تحقيق كفاءة الإنفاق في الرعاية الصحية، بإنشاء منصة متكاملة للتصنيع الحيوي على المستوى المحلي؛ بهدف الوصول إلى الاكتفاء الذاتي في إنتاج المستحضرات والمتشابهات الحيوية، وسيوفر ذلك فرصًا في لزيادة الصادرات النوعية، وتمكين نمو قطاع التقنية الحيوية.

    الجينوم:

    تُؤسس التقنية قاعدة معرفية لفهم الجينوم المحلي، وتحفيز سبل الوقاية الصحية، وتعزيز الابتكار، وذلك بتوسيع قاعدة بيانات الجينوم الوطنية ومنصة التحليلات، وتطوير السياسات الداعمة التي تُيسّر الوصول إلى البيانات، والاستفادة من هذا العلم في قطاع الرعاية الصحية؛ للحصول على تشخيصات دقيقة وفعالة.

    تحسين زراعة النباتات:

    تُحقق التقنية الاستدامة في الإمدادات الغذائية، بتحفيز الإنتاج الزراعي المحلي الذي يُحقق أثرًا إيجابيًا بتخفيض الواردات الزراعية بما فائضًا على ميزان المدفوعات، إلى جانب تعزيز الممارسات الخضراء التي تُسهم في استدامة البيئة.
',
1,
1,
1,
'2024-01-22 16:48:48',
'[{"name": "Facebook", "url": "https://facebook.com/example"},{"name": "Twitter", "url": "https://twitter.com/example"}]'
)  ,
(
  ' اقتصاد السعودية ينمو 1.2% بالربع الثاني بدعم من هذه الأنشطة ', 
  'initial/blog_3.webp', 
  array[  'الاكثر قراءة' ]::varchar[],
  'سجل الناتج المحلي الإجمالي الحقيقي في السعودية خلال الربع الثاني من عام 2023 ارتفاعًا بنسبة 1.2 بالمئة، مقارنةً بما كان عليه الفترة نفسها من العام السابق 2022، وفقًا لتقديرات الهيئة العامة للإحصاء.',
  '

سجل الناتج المحلي الإجمالي الحقيقي في السعودية خلال الربع الثاني من عام 2023 ارتفاعًا بنسبة 1.2 بالمئة، مقارنةً بما كان عليه الفترة نفسها من العام السابق 2022، وفقًا لتقديرات الهيئة العامة للإحصاء.

وقالت الهيئة في تقريرها، إن الأنشطة غير النفطية حققت خلال الربع الثاني من عام 2023 نموًا إيجابيًا بنسبة 6.1 بالمئة، مقارنةً بما كان عليه في الفترة نفسها من العام السابق 2022.

وأضافت أن الأنشطة الحكومية سجلت ارتفاعًا بنسبة 2.3 بالمئة، مقارنةً بما كان عليه في الفترة نفسها من العام السابق، فيما شهِدت الأنشطة النفطية خلال الربع الثاني من 2023، انخفاضًا بلغت نسبته 4.3 بالمئة، مقارنةً بما كان عليه في الفترة نفسها من العام السابق.

وأوضحت نتائج النشرة أن الناتج المحلي الإجمالي الحقيقي المعدَّل موسميًا شهِد خلال الربع الثاني من عام 2023 انخفاضًا بلغت نسبته 0.2 بالمئة، مقارنةً بما كان عليه في الربع الأول 2023.

وكانت التقديرات الرسمية السابقة قد توقعت نمو الناتج المحلي الإجمالي 1.1 المئة في الربع الثاني.

وسجل الاقتصاد السعودي نموا 8.7 بالمئة العام الماضي مما ساعده على تسجيل أول فائض في الميزانية منذ ما يقرب من عقد. لكن تخفيضات إنتاج النفط هذا العام وتراجع الأسعار أثرا سلبا على إيرادات النفط وسيؤثران على النمو وفي تصريحات خاصة لـ سكاي نيوز عربية، توقع رئيس بعثة صندوق النقد الدولي إلى السعودية، أمين ماتي، أن يحافظ الاقتصاد غير النفطي في السعودية، على نموه القوي هذا العام عند مستوى 4.9 بالمئة، مشيرا إلى أن بيانات معظم القطاعات تعكس متانة القطاع غير النفطي في المملكة.

توقع صندوق النقد الدولي، أن يحافظ الاقتصاد غير النفطي في السعودية، على زخم نموه القوي، بمتوسط 4.9 بالمئة، خلال العام الجاري، بفضل الطلب المحلي، وألا يتأثر النشاط غير النفطي، بتراجع النمو في إجمالي الناتج المحلي النفطي نتيجة قرارات خفض الإنتاج الطوعي.

كما يتوقع الصندوق نمو الاقتصاد السعودي في عام 2024، بنسبة 2.8 بالمئة، وأن يرتفع النمو إلى 4.2 بالمئة في 2025، و3.3 بالمئة في عامي 2026 و2027، ثم 3.1 بالمئة في عام 2028. وأشار صندوق النقد الدولي إلى أن السعودية تمكنت من احتواء التضخم عند 2.8 بالمئة في مايو 2023، كما أن البطالة وصلت إلى مستوى قياسي منخفض يبلغ 5.6 بالمئة، وسجلت مشاركة الإناث في القوى العاملة أعلى مستوياتها عند 36 بالمئة، متجاوزة الهدف المحدد عند 30 بالمئة.

يذكر أن الاقتصاد السعودي، البالغ حجمه تريليون دولار، شهد ازدهارا كبيرا في 2022، وذلك بسبب الإنتاج القياسي من الخام والذي ناهز 10.5 مليون برميل يوميا، بمتوسط سعر 100 دولار للبرميل، لكن هذا الأمر تغير في العام الجاري، بالتزامن مع تخفيضات الإنتاج التي أعلنتها أوبك+ في ظل تعثر انتعاش أسواق الطاقة العالمية. ',
1,
1,
1,
'2024-01-22 16:48:48',
'[{"name": "Facebook", "url": "https://facebook.com/example"},{"name": "Twitter", "url": "https://twitter.com/example"}]'
);

 
insert into projects(
  project_name,
  project_image,
  category_id
) values (
  'صناعة المبادرات',
  'initial/projects_initiative.webp',
  1
),
(
  'مصنع البالونات',
  'initial/projects_baloon.webp',
  2
),
(
  'الاستثمار التجاري ',
  'initial/projects_invest.webp',
  3
),
(
  'متجر المكتبة الرقمية',
  'initial/projects_ebook.webp',
  4
),
(
  'تصميم المنتجات',
  'initial/projects_prdesign.webp',
  1
);